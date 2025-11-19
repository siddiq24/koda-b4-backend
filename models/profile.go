package models

import (
	"context"
	"database/sql"
	"fmt"
)

type Profile struct {
	Id       int            `json:"id"`
	UserId   int            `json:"user_id"`
	Fullname string         `json:"full_name"`
	Image    sql.NullString `json:"image"`
	Phone    sql.NullString `json:"phone"`
	Email    string         `json:"email"`
	Address  sql.NullString `json:"address"`
}

func EditProfile(c context.Context, p Profile) (Profile, error) {
	fmt.Println("request:", p)
	if p.Fullname != "" {
		Pg.Exec(c, `UPDATE profiles SET fullname = $1 WHERE user_id = $2`, p.Fullname, p.UserId)
	}
	if p.Phone.Valid {
		Pg.Exec(c, `UPDATE profiles SET phone = $1 WHERE user_id = $2`, p.Phone, p.UserId)
	}
	if p.Email != "" {
		Pg.Exec(c, `UPDATE profiles SET email = $1 WHERE user_id = $2`, p.Phone, p.UserId)
	}
	if p.Address.Valid {
		Pg.Exec(c, `UPDATE profiles SET address = $1 WHERE user_id = $2`, p.Address, p.UserId)
	}

	if err := Pg.QueryRow(c, `SELECT fullname, image, phone, address FROM profiles WHERE user_id = $1`, p.UserId).Scan(&p.Fullname, &p.Image, &p.Phone, &p.Address); err != nil {
		return Profile{}, err
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
