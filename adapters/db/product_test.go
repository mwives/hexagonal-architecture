package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/mwives/hexagonal-architecture/adapters/db"
	"github.com/mwives/hexagonal-architecture/app"
	"github.com/stretchr/testify/assert"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	insertProduct(DB)
}

func createTable(db *sql.DB) {
	createTableStatement := `
		CREATE TABLE products (
			id STRING,
			name STRING,
			price FLOAT,
			status STRING
		);
	`

	stmt, err := db.Prepare(createTableStatement)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func insertProduct(db *sql.DB) {
	insertStatement := `
		INSERT INTO products (id, name, price, status)
		VALUES ('1', 'Product 1', 0, 'disabled');
	`

	stmt, err := db.Prepare(insertStatement)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product, err := productDB.Get("1")

	assert.Nil(t, err)
	assert.Equal(t, "1", product.GetID())
	assert.Equal(t, "Product 1", product.GetName())
	assert.Equal(t, 0.0, product.GetPrice())
	assert.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_SaveCreateProduct(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product := app.NewProduct()
	product.Name = "Product 2"
	product.Price = 10.0

	result, err := productDB.Save(product)

	assert.Nil(t, err)
	assert.Equal(t, product.GetID(), result.GetID())
	assert.Equal(t, product.GetName(), result.GetName())
	assert.Equal(t, product.GetPrice(), result.GetPrice())
	assert.Equal(t, product.GetStatus(), result.GetStatus())
}

func TestProductDB_SaveUpdateProduct(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product, _ := productDB.Get("1")
	product.Enable()

	result, err := productDB.Save(product)

	assert.Nil(t, err)
	assert.Equal(t, product.GetID(), result.GetID())
	assert.Equal(t, product.GetName(), result.GetName())
	assert.Equal(t, product.GetPrice(), result.GetPrice())
	assert.Equal(t, product.GetStatus(), result.GetStatus())
}
