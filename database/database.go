package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error){
	var err error
	dns := "user=postgres password=postgres dbname=app-destinations sslmode=disable host=localhost"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database; %v", err)
	}

	fmt.Println("Database connection successfully!")
	return db, nil
}
