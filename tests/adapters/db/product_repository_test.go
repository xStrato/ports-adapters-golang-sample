package db_test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/ports-adapters-go-sample/src/adapters/db"
	"github.com/xStrato/ports-adapters-go-sample/src/application/constants"
)

var Db *sql.DB

func TestProductRepository(t *testing.T) {
	setUp()

	t.Run("Get_ValidProduct_ShouldReturnNil", func(t *testing.T) {
		//Arrange
		productDB := db.NewProductRepository(Db)
		//Act
		product, err := productDB.Get("abx")
		//Assert
		require.Nil(t, err)
		require.Equal(t, "Test", product.GetName())
		require.Equal(t, 0.0, product.GetPrice())
		require.Equal(t, constants.DISABLED, product.GetStatus())
	})
	Db.Close()
}

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProducts(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE [products](
		[id] TEXT,
		[name] TEXT,
		[price] FLOAT,
		[status] TEXT
		);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatalln(err.Error())
	}
	stmt.Exec()
}

func createProducts(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abx", "Test", 0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	stmt.Exec()
}
