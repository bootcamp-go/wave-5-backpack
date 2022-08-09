package domain

type Transanction struct {
	Id       int     `form:"id" json:"id"`
	Code     string  `form:"code" json:"code" binding:"required"`
	Coin     string  `form:"coin" json:"coin" binding:"required"`
	Amount   float64 `form:"amount" json:"amount" binding:"required"`
	Emisor   string  `form:"emisor" json:"emisor" binding:"required"`
	Receptor string  `form:"receptor" json:"receptor" binding:"required"`
	Date     string  `form:"date" json:"date" binding:"required"`
}
