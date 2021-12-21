package entity

import "strconv"

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

func (product Product) StockStatus() (status string) {
	if product.Stock == 0 {
		status = "Out of stock"
	} else if product.Stock < 3 {
		status = "Stock is running out"
	} else if product.Stock < 10 {
		status = "Stock is limited"
	} else {
		status = "Stock: " + strconv.Itoa(product.Stock)
	}

	return
}
