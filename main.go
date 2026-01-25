package main

import (
	//"bufio"
	//"os"
	//"strings"
	//"fmt"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	pool := connect()
	var wrestler Wrestler
	wrestler.Ringname = "Kurt Angle"
	wrestler.Alignment = "face"
	wrestler.SignatureMove = "Angle Slam"

	repo := wrestlerRepository{
		dbpool: pool,
	}
	repo.Create(&wrestler)
}