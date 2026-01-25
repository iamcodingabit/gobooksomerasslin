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

	query := `INSERT INTO wrestlers(ringname, alignment, signature_move) 
		VALUES ($1, $2, $3) 
		RETURNING wrestler_id;`
	
	err := w.dbpool.QueryRow(c, query, wrestler.Ringname, wrestler.Alignment, wrestler.SignatureMove).Scan(&id)
	if err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}

	fmt.Printf("Created wrestler with ID: %d\n", id)
	return nil
}

func (w *wrestlerRepository) ReadWrestlerByRingname(c context.Context, ringname string) (Wrestler, error){
	var wrestler Wrestler
	query := "SELECT ringname, alignment, signature_move FROM wrestlers where ringname LIKE $1"

	err := w.dbpool.QueryRow(c, query, ringname).Scan(&wrestler.Ringname, &wrestler.Alignment, &wrestler.SignatureMove)
	if err != nil {
		return wrestler, fmt.Errorf("error reading task: %w", err)
	}

	return wrestler, nil
}

func (w *wrestlerRepository) ReadAllWrestler(c context.Context) ([]Wrestler,  error) {
	var wrestlers []Wrestler
	query := "SELECT ringname, alignment, signature_move FROM wrestlers ORDER BY ringname"

	rows, err := w.dbpool.Query(c, query)
	if err != nil {
		return wrestlers, fmt.Errorf("error querying wreslters: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var wrestler Wrestler
		err := rows.Scan(
			&wrestler.Ringname,
			&wrestler.Alignment,
			&wrestler.SignatureMove,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning wrestler row: %w", err)
		}
		wrestlers = append(wrestlers, wrestler)
	}

	return wrestlers, nil
}