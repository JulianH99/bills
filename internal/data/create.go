package data

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (d database) createTables() error {
	for _, table := range tables {
		_, err := d.instance.Exec(table.create)
		if err != nil {
			fmt.Println("Error creating tarble", table.name)
			return err
		}
	}
	return nil
}
