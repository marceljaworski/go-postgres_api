package models

type Product struct {
	ProductID int64  `json:productid`
	Name      string `json:name`
	Price     int64  `json:price`
	Company   string `json:company`
}
