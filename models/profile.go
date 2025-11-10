package models

import "context"

type Profile struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Fullname string `json:"full_name"`
	Image    string `json:"image"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func EditProfile(c context.Context, p Profile) (Profile, error) {
	if p.Fullname != "" {
		Pg.Exec(c, `UPDATE profiles SET fullname = $1 WHERE user_id = $2`, p.Fullname, p.UserId)
	}
	if p.Image != "" {
		Pg.Exec(c, `UPDATE profiles SET image = $1 WHERE user_id = $2`, p.Image, p.UserId)
	}
	if p.Phone != "" {
		Pg.Exec(c, `UPDATE profiles SET phone = $1 WHERE user_id = $2`, p.Phone, p.UserId)
	}
	if p.Address != "" {
		Pg.Exec(c, `UPDATE profiles SET address = $1 WHERE user_id = $2`, p.Address, p.UserId)
	}

	if err := Pg.QueryRow(c, `SELECT fullname, image, phone, address FROM profiles WHERE user_id = $1`, p.UserId).Scan(&p.Fullname, &p.Image, &p.Phone, &p.Address); err != nil {
		return Profile{}, err
	}
	return p, nil
}
