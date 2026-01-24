package main

import (
	//"database/sql"
	//"github.com/lib/pq"
	"context"
	"os"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func connect(){
	err := godotenv.Load()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var wrestler Wrestler
	err = conn.QueryRow(context.Background(), "select ringname, alignment, signature_move FROM wrestlers where ringname LIKE $1;","%King%Mystery%").Scan(&wrestler.Ringname, &wrestler.Alignment, &wrestler.SignatureMove)


	fmt.Printf("%s - %s - %s\n", wrestler.Ringname, wrestler.Alignment, wrestler.SignatureMove)
}

