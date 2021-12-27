package interfaces

type Entity interface {
	GetID() string
	IsValid() (bool, error)
	Enable() error
	Disable() error
}
