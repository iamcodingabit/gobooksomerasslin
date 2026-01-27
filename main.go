package main

import (
	//"bufio"
	//"os"
	//"strings"
	//"context"
	"fmt"
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
	
	wrestlerList := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true)
	wrestlerList.SetBorder(true).SetTitle("Wrestlers")

	for _, wrestler := range wrestlers {
		fmt.Fprintf(wrestlerList, "%s - %s - %s\n", wrestler.Ringname, wrestler.Alignment, wrestler.SignatureMove)
	}

	flex := tview.NewFlex().AddItem(wrestlerList, 0, 1, false)

	if err:=app.SetRoot(flex, true).Run(); err != nil{
		panic(err)
	}
	
	defer pool.Close()
}