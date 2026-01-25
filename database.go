package main

import (
	//"database/sql"
	//"github.com/lib/pq"
	"context"
	"fmt"
	"log"
	"os"

	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func connect()*pgxpool.Pool{
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()) 
	err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	fmt.Println("Connected to database!")
	return dbpool
}

