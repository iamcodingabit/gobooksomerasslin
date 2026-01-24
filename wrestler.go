package main

type Wrestler struct {
	ringname string `json:"ringname"`
	alignment string `json:"alignment"` // face | heel | tweener
	signature_move string `json:"signature_move"`
}

var wrestlers = []Wrestler{
	{ringname: "King Mystery", alignment: "face", signature_move: "Area Codes"},
	{ringname: "Jonathan C. Nah", alignment: "face", signature_move: "Fix Your Face"},
	{ringname: "HANS", alignment: "heel", signature_move: "Piledriver"},
}