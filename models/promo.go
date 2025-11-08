package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Promo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func AllPromo(c context.Context, pg *pgxpool.Pool) ([]Promo, error) {
	var promos []Promo
	Query := `SELECT id, title, description FROM promos WHERE "end" >= CURRENT_DATE`
	rows, err := pg.Query(c, Query)
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
