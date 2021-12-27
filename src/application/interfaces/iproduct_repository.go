package interfaces

type IProductRepository interface {
	Get(id string) (IProduct, error)
	Create(name string, price float32) (IProduct, error)
}
