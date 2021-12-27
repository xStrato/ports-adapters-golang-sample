package application

import (
	"errors"

	"github.com/xStrato/ports-adapters-go-sample/src/application/constants"
)

type Product struct {
	id     string
	name   string
	price  float32
	status constants.Status
}

func NewProduct(name string, price float32) *Product {
	return &Product{
		id:     "12334",
		name:   name,
		price:  price,
		status: constants.DISABLE,
	}
}

func (p *Product) IsValid() bool {
	return false
}

func (p *Product) Enable() error {
	if p.price > 0 {
		p.status = constants.ENABLE
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

func (p *Product) GetID() string {
	return p.id
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetStatus() string {
	return string(p.status)
}

func (p *Product) GetPrice() float32 {
	return p.price
}
