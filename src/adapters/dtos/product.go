package dtos

import (
	"github.com/xStrato/ports-adapters-go-sample/src/application/entities"
	"github.com/xStrato/ports-adapters-go-sample/src/application/interfaces"
)

type Product struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct(id, name, status string, price float32) *Product {
	return &Product{id, name, price, status}
}
func (p *Product) Bind(product interfaces.IProduct) (interfaces.IProduct, error) {
	binded := entities.NewProductFrom(p.Id, p.Name, p.Status, p.Price)
	_, err := binded.IsValid()
	if err != nil {
		return nil, err
	}
	return binded, nil
}
