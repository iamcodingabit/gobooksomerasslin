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
		ReadAllWrestler() ([]Wrestler,  error)
		ReadWrestlerByRingname(ringname string) (Wrestler, error)
		Update(ringname string) (Wrestler, error)
		Delete(ringname string) (Wrestler, error)
	}
)