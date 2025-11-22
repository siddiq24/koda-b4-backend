package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type ProfileRequest struct {
	UserId    int       `json:"userId"`
	Fullname  string    `json:"fullName"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Profile struct {
	Id        int            `json:"id"`
	UserId    int            `json:"userId"`
	Fullname  string         `json:"fullname"`
	Image     sql.NullString `json:"image"`
	Phone     sql.NullString `json:"phone"`
	Email     string         `json:"email"`
	Address   sql.NullString `json:"address"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt sql.NullTime   `json:"updatedAt"`
}

func EditProfile(c context.Context, p ProfileRequest) (ProfileRequest, error) {
	var (
		existingFullname string
		existingPhone    sql.NullString
		existingAddress  sql.NullString
	)

	err := Pg.QueryRow(c, `
		SELECT fullname, phone, address
		FROM profiles 
		WHERE user_id = $1
	`, p.UserId).Scan(&existingFullname, &existingPhone, &existingAddress)

	if err != nil {
		return ProfileRequest{}, fmt.Errorf("failed to get existing profile: %w", err)
	}

	if p.Fullname == "" {
		p.Fullname = existingFullname
	}
	if p.Phone == "" {
		p.Phone = existingPhone.String
	}
	if p.Address == "" {
		p.Address = existingAddress.String
	}

	err = Pg.QueryRow(c, `
		UPDATE profiles
		SET fullname = $1, phone = $2, address = $3
		WHERE user_id = $4
		RETURNING fullname, phone, address
	`, p.Fullname, p.Phone, p.Address, p.UserId).Scan(
		&p.Fullname,
		&p.Phone,
		&p.Address,
	)

	if err != nil {
		return ProfileRequest{}, fmt.Errorf("failed to update profile: %w", err)
	}

	return p, nil
}

func UpdateImage(c context.Context, Uid int, image string) (string, error) {
	var img string
	if image != "" {
		if err := Pg.QueryRow(c, `UPDATE profiles SET image = $1 WHERE user_id = $2 RETURNING image`, image, Uid).Scan(&img); err != nil {
			return "", err
		}
	}

	return img, nil
}

func GetProfileInfo(c context.Context, Uid int) (Profile, error) {
	var profile Profile
	if err := Pg.QueryRow(c, `
		SELECT fullname, image, phone, address 
		FROM profiles 
		WHERE user_id = $1`,
		Uid).
		Scan(
			&profile.Fullname,
			&profile.Image,
			&profile.Phone,
			&profile.Address,
		); err != nil {
		return Profile{}, err
	}
	if err := Pg.QueryRow(c, `
		SELECT email 
		FROM users 
		WHERE id = $1`,
		Uid).
		Scan(
			&profile.Email,
		); err != nil {
		return Profile{}, err
	}

	return profile, nil
}
