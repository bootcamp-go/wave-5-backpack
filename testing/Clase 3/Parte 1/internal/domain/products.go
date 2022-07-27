package domain

type Product struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      uint64  `json:"stock"`
	Code       string  `json:"code"`
	Published  bool    `json:"published"`
	Created_at string  `json:"created_at"`
}
