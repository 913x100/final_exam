package database

import (
	"database/sql"
	"fmt"
)

func (d *Database)GetCustomers() (*sql.Rows, error) {
	sql, err := d.DB.Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		return nil, fmt.Errorf("can't prepare select statement: %w", err)
	}

	rows, err := sql.Query()
	if err != nil {
		return nil, fmt.Errorf("can't get customers data: %w", err)
	}

	return rows, nil
}

func (d *Database)GetCustomer(id string) (*sql.Row, error) {
	sql, err := d.DB.Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		return nil, fmt.Errorf("can't prepare select statement: %w", err)
	}

	row := sql.QueryRow(id)
	return row, nil
}