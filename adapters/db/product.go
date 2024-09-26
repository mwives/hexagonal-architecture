package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mwives/hexagonal-architecture/app"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db}
}

func (p *ProductDB) Get(id string) (app.ProductInterface, error) {
	var product app.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Save(product app.ProductInterface) (app.ProductInterface, error) {
	var rows int
	p.db.QueryRow(
		"SELECT COUNT(*) FROM products WHERE id = ?",
		product.GetID(),
	).Scan(&rows)

	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return &app.Product{}, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return &app.Product{}, err

		}
	}
	return product, nil
}

func (p *ProductDB) create(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare(
		"INSERT INTO products(id, name, price, status) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) update(product app.ProductInterface) (app.ProductInterface, error) {
	_, err := p.db.Exec(
		"UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
