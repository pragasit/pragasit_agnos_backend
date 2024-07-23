package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func connectDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@db:5432/mydb")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func logRequestResponse(request string, response string) error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "INSERT INTO logs (request, response) VALUES ($1, $2)", request, response)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := logRequestResponse("test request", "test response")
	if err != nil {
		fmt.Printf("Failed to log request/response: %v\n", err)
	} else {
		fmt.Println("Successfully logged request/response")
	}
}
