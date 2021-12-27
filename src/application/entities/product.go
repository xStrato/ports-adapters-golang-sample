package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"github.com/xStrato/ports-adapters-go-sample/src/application/constants"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	id     string           `valid:"uuidv4"`
	name   string           `valid:"required"`
	price  float32          `valid:"float,optional"`
	status constants.Status `valid:"required"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{
		id:     uuid.NewV4().String(),
		name:   name,
		price:  price,
		status: constants.DISABLED,
	}
}

func (p *Product) IsValid() (bool, error) {
	if p.status == "" {
		p.status = constants.DISABLED
	}
	if p.status != constants.ENABLED && p.status != constants.DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}
	if p.price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.price > 0 {
		p.status = constants.ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.price == 0 {
		p.status = constants.DISABLED
		return nil
	}
	return errors.New("the price must be zero in order to have the product disabled")
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
