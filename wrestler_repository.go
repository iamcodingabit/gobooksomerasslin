package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type wrestlerRepository struct {
	dbpool *pgxpool.Pool
}

func (w *wrestlerRepository) Create(c context.Context, wrestler *Wrestler) error{
	var id int
	fmt.Println(wrestler.Ringname)
	fmt.Println(wrestler.Alignment)
	fmt.Println(wrestler.SignatureMove)

	query := `
		INSERT INTO wrestlers(ringname, alignment, signature_move) 
		VALUES ($1, $2, $3) 
		RETURNING wrestler_id;
	`
	
	err := w.dbpool.QueryRow(c, query, wrestler.Ringname, wrestler.Alignment, wrestler.SignatureMove).Scan(&id)
	if err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}

	fmt.Printf("Created wrestler with ID: %d\n", id)
	return nil

	/*
	var wrestler Wrestler
	err = dbpool.QueryRow(context.Background(), "select ringname, alignment, signature_move FROM wrestlers where ringname LIKE $1;","%King%Mystery%").Scan(&wrestler.Ringname, &wrestler.Alignment, &wrestler.SignatureMove)


	fmt.Printf("%s - %s - %s\n", wrestler.Ringname, wrestler.Alignment, wrestler.SignatureMove)
	*/
}