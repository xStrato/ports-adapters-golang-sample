package interfaces

type Entity interface {
	IsValid() bool
	Enable() error
	Disable() error
}
