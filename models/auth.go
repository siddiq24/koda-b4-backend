package models

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthRequest struct {
	Fullname string `json:"fullname" form:"fullname" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"min=8"`
	Role     string `json:"role"`
}

type AuthResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Auth struct{}

func (m *Auth) AddUser(c context.Context, newUser AuthRequest) (int, error) {
	var ID int
	User_sql := `INSERT INTO users(email, password, role) VALUES ($1, $2, $3) RETURNING id;`
	if err := Pg.QueryRow(c, User_sql, newUser.Email, newUser.Password, newUser.Role).Scan(&ID); err != nil {
		log.Println(err)
		return 0, fmt.Errorf("failed insert user")
	}
	Profile_sql := `INSERT INTO profiles(fullname, user_id) VALUES ($1, $2)`
	Pg.QueryRow(c, Profile_sql, newUser.Fullname, ID)
	return ID, nil
}

func (m *Auth) EmailExist(c context.Context, email string) int {
	ID := 0
	User_sql := `SELECT id FROM users WHERE email = $1`
	if err := Pg.QueryRow(c, User_sql, email).Scan(&ID); err != nil {
		log.Println("Email not exist")
		return -1
	}
	return ID
}

func (m *Auth) PasswordIDUser(c context.Context, email string) (int, string, string, error) {
	var id int
	var pass string
	var role string

	query := `SELECT id, password, role FROM users WHERE email = $1`

	err := Pg.QueryRow(c, query, email).Scan(&id, &pass, &role)
	if err != nil {
		log.Println("Query error:", err)
		return 0, "", "", err
	}

	return id, pass, role, nil
}

func (m *Auth) ForgotPassword(ctx context.Context, email string) (string, error) {
	tokenInt, err := rand.Int(rand.Reader, big.NewInt(100000000))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	token := fmt.Sprintf("%08d", tokenInt.Int64())

	ttl := time.Minute * 2

	val, err := Rdb.Get(ctx, email).Result()
	if err == nil && val != "" {
		return "", fmt.Errorf("token sudah dikirim, coba beberapa saat lagi")
	}

	if err != nil && err != redis.Nil {
		return "", fmt.Errorf("redis error: %v", err)
	}

	if err := Rdb.Set(ctx, email, token, ttl).Err(); err != nil {
		return "", fmt.Errorf("failed to save token: %v", err)
	}

	return token, nil
}

func (m *Auth) ValidatePIN(ctx context.Context, email, pin string) bool {
	val, err := Rdb.Get(ctx, email).Result()
	if err == redis.Nil {
		fmt.Println("token tidak ada atau expired")
		return false
	}
	if err != nil {
		log.Println("redis error:", err)
		return false
	}

	return val == pin
}

type ForgotPassword struct {
	Email       string `json:"email"`
	Pin         string `json:"pin"`
	NewPassword string `json:"new_password"`
	Origin      string `json:"origin"`
}

func (m *Auth) UpdatePassword(ctx context.Context, req ForgotPassword) error {
	_, err := Pg.Exec(ctx, `
		UPDATE users
		SET password = $1
		WHERE email = $2
	`, req.NewPassword, req.Email)

	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	_ = Rdb.Del(ctx, req.Email).Err()

	return nil
}

func (m *Auth) Logout(ctx context.Context, tokenString string, exp time.Duration) error {
	if tokenString == "" {
		return errors.New("token kosong")
	}

	key := "jwt:blacklist:" + tokenString[7:]
	fmt.Println("key from model:", key)

	err := Rdb.Set(ctx, key, "1", exp).Err()
	if err != nil {
		return fmt.Errorf("gagal mem-blacklist token: %w", err)
	}

	return nil
}
