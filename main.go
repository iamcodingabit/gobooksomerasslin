package main

import (
	//"bufio"
	//"os"
	//"strings"
	"fmt"
	"context"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	pool := connect()

	w := wrestlerRepository {
		dbpool: pool,		
	}

	wrestlers, _ := w.ReadAllWrestler(context.Background())

	for _, v := range(wrestlers){
		fmt.Printf("%s - %s - %s\n", v.Ringname, v.Alignment, v.SignatureMove)
	}

	defer pool.Close()
}