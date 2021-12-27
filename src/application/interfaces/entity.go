package interfaces

type Entity interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
}
