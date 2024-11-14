package main

import (
	"fmt"

	"github.com/shivkr6/go-rest-api-comments/internal/db"
)

func Run() error {
	fmt.Println("startup up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go REST API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
