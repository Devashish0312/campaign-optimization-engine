package db

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func ConnectToDatabase() error {
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres dbname=campaigns sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
		return err
	}
	fmt.Println("Connected to Database")
	return nil
}
