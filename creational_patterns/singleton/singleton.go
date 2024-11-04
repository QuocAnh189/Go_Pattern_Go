package database

import (
	"fmt"
	"sync"
)

type Database struct {
	connectionString string
}

func (db *Database) Query(query string) {
	fmt.Println("Executing query:", query)
}

var (
	instance *Database
	once     sync.Once
)

func GetDatabaseInstance(connectionString string) *Database {
	once.Do(func() {
		fmt.Println("Creating database connection...")
		instance = &Database{connectionString: connectionString}
	})
	return instance
}
