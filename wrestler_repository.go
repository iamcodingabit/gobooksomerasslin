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

func (w *wrestlerRepository) ReadAllWrestler() (error) {
	var wrestlers []Wrestler
	query := "SELECT ringname, alignment, signature_move FROM wrestlers ORDER BY ringname"
	c := context.Background()
	rows, err := w.dbpool.Query(c, query)
	if err != nil {
		return fmt.Errorf("error querying wreslters: %w", err)
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
			return fmt.Errorf("error scanning wrestler row: %w", err)
		}
		wrestlers = append(wrestlers, wrestler)
	}

	for _, v := range(wrestlers){
		fmt.Printf("%s - %s - %s\n", v.Ringname, v.Alignment, v.SignatureMove)
	}

	return nil
}

func (w *wrestlerRepository) UpdateRingname(c context.Context, current_ringname string, new_ringname string) (error){
	query := `
		UPDATE wrestlers
		SET ringname = $2
		WHERE ringname LIKE $1
		RETURNING $2;
	`
	err := w.dbpool.QueryRow(c, query, current_ringname, new_ringname).Scan(&new_ringname)
	if (err!=nil){
		return fmt.Errorf("error updating wrestler row: %w", err)
	}

	return nil
}

func (w *wrestlerRepository) DeleteByRingname(c context.Context, ringname string) (string, error){
	var deleted_ringname string
	query := `
		DELETE FROM wrestlers WHERE ringname LIKE $1 RETURNING ringname;
	`
	err := w.dbpool.QueryRow(c, query, ringname).Scan(&deleted_ringname)
	if err != nil {
		return "", fmt.Errorf("error deleting wrestler row: %w", err)
	}

	return ringname, nil
}