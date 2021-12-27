package interfaces

type IProductService interface {
	Get(id string) (IProduct, error)
	Create(name string, price float32) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error)
}
