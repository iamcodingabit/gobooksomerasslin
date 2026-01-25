package main

import "context"

type Wrestler struct {
	Ringname string `json:"ringname"`
	Alignment string `json:"alignment"` // face | heel | tweener
	SignatureMove string `json:"signature_move"`
}

type (
	WrestlerRepository interface{
		Create(c context.Context, wrestler *Wrestler) error
		ReadAllWrestler(c context.Context) ([]Wrestler,  error)
		ReadWrestlerByRingname(c context.Context, ringname string) (Wrestler, error)
		UpdateRingname(c context.Context, current_ringname string, new_ringname string) (Wrestler, error)
		Delete(c context.Context, ringname string) (Wrestler, error)
	}
)