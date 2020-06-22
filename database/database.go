package database

import (
	"database/sql"
	"final_exam/models"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type IDatabase interface {
	CreateCustomer(cus models.Customer) (int, error)
	GetCustomers() (*sql.Rows, error)
	GetCustomer(id string) (*sql.Row, error)
	UpdateCustomer(cus models.Customer) error
	DeleteCustomer(id string) error
}

type Database struct {
	DB *sql.DB
}

func New() *Database {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	createCustomerTable(db)

	return &Database{DB: db}
}

func createCustomerTable(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);`

	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}
}
