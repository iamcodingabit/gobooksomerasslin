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

	fmt.Println(w.ReadWrestlerByRingname(context.Background(), "Kurt Angle"))

	defer pool.Close()
}