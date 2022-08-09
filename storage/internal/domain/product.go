package domain

type Product struct {
	ID     int     `db:"ID"`
	Name   string  `db:"name"`
	Type   string  `db:"type"`
	Price  float64 `db:"price"`
	Count  int     `db:"count"`
	Code   string  `db:"code"`
	Public int8    `db:"public"`
}
