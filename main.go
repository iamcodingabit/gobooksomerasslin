package main

import (
	//"bufio"
	//"os"
	//"strings"
	//"fmt"
	//"context"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	pool := connect()
	defer pool.Close()
}