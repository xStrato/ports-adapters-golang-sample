package services

import "github.com/xStrato/ports-adapters-go-sample/src/application/interfaces"

type ProductService struct {
	repository interfaces.IProductRepository
}

func NewProductService(r interfaces.IProductRepository) *ProductService {
	return &ProductService{
		repository: r,
	}
}

func (s *ProductService) Get(id string) (interfaces.IProduct, error) {
	product, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float32) (interfaces.IProduct, error) {
	product, err := s.repository.Create(name, price)
	if err != nil {
		return nil, err
	}
	return product, nil
}
