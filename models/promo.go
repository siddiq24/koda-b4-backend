package models

import (
	"context"
)

type Promo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func AllPromo(c context.Context) ([]Promo, error) {
	var promos []Promo
	Query := `SELECT id, title, description FROM promos WHERE "end" >= CURRENT_DATE`
	rows, err := Pg.Query(c, Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Promo
		if err := rows.Scan(&p.Id, &p.Title, &p.Desc); err != nil {
			return nil, err
		}
		promos = append(promos, p)
	}
	return promos, nil

}
