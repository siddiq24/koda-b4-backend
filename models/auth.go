package models

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
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

type Auth struct {
	Pg *pgxpool.Pool
}

func NewAuth(pg *pgxpool.Pool) *Auth {
	return &Auth{Pg: pg}
}

func (m *Auth) AddUser(c context.Context, newUser AuthRequest) (int, error) {
	var ID int
	User_sql := `INSERT INTO users(email, password, role) VALUES ($1, $2, $3) RETURNING id;`
	if err := m.Pg.QueryRow(c, User_sql, newUser.Email, newUser.Password, newUser.Role).Scan(&ID); err != nil {
		log.Println(err)
		return 0, fmt.Errorf("failed insert user")
	}
	Profile_sql := `INSERT INTO profiles(fullname, user_id) VALUES ($1, $2)`
	m.Pg.QueryRow(c, Profile_sql, newUser.Fullname, ID)
	return ID, nil
}

func (m *Auth) EmailExist(c context.Context, email string) int {
	ID := 0
	User_sql := `SELECT id FROM users WHERE email = $1`
	if err := m.Pg.QueryRow(c, User_sql, email).Scan(&ID); err != nil {
		log.Println(err)
	}
	return ID
}

func (m *Auth) PasswordIDUser(c context.Context, email string) (int, string) {
	var id int
	var pass string
	Query := `SELECT id, password FROM users WHERE email=$1`
	if err := m.Pg.QueryRow(c, Query, email).Scan(&id, &pass); err != nil {
		log.Println(err)
		return 0, ""
	}
	return id, pass
}
