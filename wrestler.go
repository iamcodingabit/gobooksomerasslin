package main

type Wrestler struct {
	Ringname string `json:"ringname"`
	Alignment string `json:"alignment"` // face | heel | tweener
	SignatureMove string `json:"signature_move"`
}

type (
	/*
	WrestlerUsecase interface{
		Sign(wrestler *Wrestler) error
		List() ([]Wrestler,  error)
		Get(ringname string) (Wrestler, error)
		Update(ringname string) (Wrestler, error)
		Release(ringname string) (Wrestler, error)
	}
	*/

	WrestlerRepository interface{
		Create(wrestler *Wrestler) error
		ReadAllWrestler() ([]Wrestler,  error)
		ReadWrestlerByRingname(ringname string) (Wrestler, error)
		Update(ringname string) (Wrestler, error)
		Delete(ringname string) (Wrestler, error)
	}
)