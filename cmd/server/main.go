package main

import (
	"context"
	"fmt"

	"github.com/shivkr6/go-rest-api-comments/internal/comment"
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

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(context.Background(), "550e8400-e29b-41d4-a716-446655440000"))
	return nil
}

func main() {
	fmt.Println("Go REST API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
