package models

type Product struct {
	ProductID int64  `json:"product_id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Company   string `json:"company"`
}
