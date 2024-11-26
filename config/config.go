package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dsn := `root:@tcp(127.0.0.1:3306)/databasepetcare`
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("1Cosssssnnected to the database!")
	return db
}
