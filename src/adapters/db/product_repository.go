package db

import (
	"database/sql"

	"github.com/xStrato/ports-adapters-go-sample/src/application/entities"
	"github.com/xStrato/ports-adapters-go-sample/src/application/interfaces"
)

type ProductRepository struct {
	*sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Get(id string) (interfaces.IProduct, error) {
	product := entities.Product{}
	stmt, err := r.DB.Prepare("select name from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
