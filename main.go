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

	wrestlerList := tview.NewList().ShowSecondaryText(false)
	for _, wrestler := range wrestlers {
		wrestlerList.AddItem(wrestler.Ringname, "Test", 0, func(){})
	}

	//textboox for insertion
	//insert_flex := tview.NewFlex().AddItem(tview.NewTextArea(), 0, 1, false)

	menuList := tview.NewList().
		AddItem("Create wrestler", "", 0, func(){
		}).
		AddItem("Remove wrestler", "", 0, func(){})
	
	/*
	flex := tview.NewFlex().
		AddItem(wrestlerList, 0, 1, false).
		AddItem(randomList, 0, 1, true)
	*/

	grid := tview.NewGrid().SetRows(3, 0, 3).SetColumns(30, 0,30).SetBorders(true)

	button_prompts := tview.NewTextView()
	fmt.Fprintf(button_prompts, "c - create wrestler")


	grid.AddItem(wrestlerList, 1, 2, 1, 1, 0, 0, false).
		AddItem(menuList , 1, 0, 1, 1, 0, 0, true).
		AddItem(button_prompts, 2, 0, 1, 3, 0, 0, false)
	
	if err:=app.SetRoot(grid, true).Run(); err != nil{
		panic(err)
	}
	
	defer pool.Close()
}