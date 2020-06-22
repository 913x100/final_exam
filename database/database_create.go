package database

import (
	"final_exam/models"
	"fmt"
)

func (d *Database) CreateCustomer(cus models.Customer) (int, error) {
	var id int
	row := d.DB.QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", cus.Name, cus.Email, cus.Status)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can't create customers: %w", err)
	}

	return id, nil
}
