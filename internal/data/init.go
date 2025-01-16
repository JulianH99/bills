package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	fp       string
	instance *sql.DB
}

func NewDatabase(fp string) database {
	return database{fp: fp}
}

// Initializes the database and returns a pointer to a DB struct
//
// fp string filepath to save the database into
func (d *database) Open() error {
	db, err := sql.Open("sqlite3", d.fp)
	if err != nil {
		return err
	}

	d.instance = db

	return nil
}

func (d *database) Close() {
	d.instance.Close()
}

func (d database) Instance() *sql.DB {
	return d.instance
}

// Initializes everything related to the database
// As per now, it only creates the tables needed for the
// application to work
func (d database) Initialize() error {
	return d.createTables()
}
