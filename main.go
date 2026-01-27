package main

import (
	//"bufio"
	//"os"
	//"strings"
	//"context"
	//"fmt"
	"github.com/iamcodingabit/gobooksomerasslin/database"
	"github.com/joho/godotenv"
	"github.com/rivo/tview"
)

func main(){
	godotenv.Load()
	pool := database.Connect()
	//c := context.Background()
	var w WrestlerRepository = &wrestlerRepository{ 
		dbpool: pool, 
	}

	app := tview.NewApplication()

	wrestlers, _ := w.ReadAllWrestler()

	wrestlerList2 := tview.NewList().ShowSecondaryText(false)

	for _, wrestler := range wrestlers {
		wrestlerList2.AddItem(wrestler.Ringname, "", 0, func(){})
	}
	
	flex := tview.NewFlex().
		AddItem(wrestlerList2, 0, 1, false)

	if err:=app.SetRoot(flex, true).Run(); err != nil{
		panic(err)
	}
	
	defer pool.Close()
}