package singleton

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


func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		db := GetDatabaseInstance("DB_CONNECTION_STRING")
		db.Query("SELECT * FROM users")
	}()

	go func() {
		defer wg.Done()
		db := GetDatabaseInstance("DB_CONNECTION_STRING")
		db.Query("SELECT * FROM orders")
	}()

	wg.Wait()

	db1 := GetDatabaseInstance("DB_CONNECTION_STRING")
	db2 := GetDatabaseInstance("DB_CONNECTION_STRING")

	if db1 == db2 {
		fmt.Println("Only one instance of Database exists.")
	} else {
		fmt.Println("Different instances exist, something went wrong!")
	}
}