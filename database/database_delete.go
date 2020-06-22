package database

import "fmt"

func (d *Database) DeleteCustomer(id string) error {
	sql, err := d.DB.Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err := sql.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete statement: %w", err)
	}

	return nil
}
