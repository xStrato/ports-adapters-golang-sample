package interfaces

import "github.com/xStrato/ports-adapters-go-sample/src/application/constants"

type IProduct interface {
	GetID() string
	GetName() string
	GetStatus() constants.Status
	GetPrice() float32
	IsValid() (bool, error)
	Enable() error
	Disable() error
}
