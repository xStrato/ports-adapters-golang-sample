package db

import (
	"database/sql"

	"github.com/blockloop/scan"
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
	stmt, err := r.DB.Prepare("select * from products where id=?")
	if err != nil {
		return nil, err
	}
	row, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	p := struct {
		Id, Name string
		Price    float32
		Status   string
	}{}
	err = scan.Row(&p, row)
	if err != nil {
		return nil, err
	}
	return entities.NewProductFrom(p.Id, p.Name, p.Status, p.Price), nil
}
func (r *ProductRepository) Save(p interfaces.IProduct) (interfaces.IProduct, error) {
	var rows int
	r.DB.QueryRow("select id from products where id=?", p.GetID()).Scan(&rows)
	if rows == 0 {
		if _, err := r.create(p); err != nil {
			return nil, err
		}
		return p, nil
	}
	if _, err := r.update(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) create(p interfaces.IProduct) (interfaces.IProduct, error) {
	stmt, err := r.DB.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.GetID(), p.GetName(), p.GetPrice(), p.GetStatus())
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) update(p interfaces.IProduct) (interfaces.IProduct, error) {
	stmt, err := r.DB.Prepare(`update products set name=?, price=?, status=? where id=?;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.GetName(), p.GetPrice(), p.GetStatus(), p.GetID())
	if err != nil {
		return nil, err
	}
	return p, nil
}
