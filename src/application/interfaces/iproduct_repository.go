package interfaces

type IProductRepository interface {
	Get(id string) (IProduct, error)
	Save(product IProduct) (IProduct, error)
}
