package storage

import (
	"github.com/asdine/storm/v3"
)

var db *storm.DB

func GetDB() *storm.DB {
	return db
}

// Initialize opens and set the database based on the given path.
func Initialize(path string) error {
	d, err := storm.Open(path)
	db = d
	return err
}

func Close() {
	db.Close()
}
