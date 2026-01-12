package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
)

type Wrestler struct {
	ringname string
	nickname string
	alignment string
	signature string
	finisher string
}

var Commands = map[string]interface{}{
	"add": addWrestler,
	"remove": addWrestler,
	"edit": addWrestler,
	"finish": addWrestler,
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	wrestlers := []Wrestler{}
	
	input := ""
	isActive := true
	
	for isActive {
		fmt.Print("Add wrestlers by typing 'add', type 'quit' to quit: ")
		scanner.Scan()
		
		input = scanner.Text()
		command := strings.SplitAfterN(input, " ", 2)

		if _, ok := Commands[strings.Trim(command[0], " ")]; ok {
			Commands[strings.Trim(command[0], " ")].(func(*[]Wrestler, string))(&wrestlers, command[1])
			listWrestlers(wrestlers)
		} else {
			isActive = false
		}
	}
	
	fmt.Println("Thanks for using the app!")
}

func addWrestler(wrestlers *[]Wrestler, newWrestler string){
	var var1 Wrestler
	var1.ringname = newWrestler
	*wrestlers = append(*wrestlers, var1)
}

func listWrestlers(wrestlers []Wrestler){
	for _, v := range wrestlers {
		fmt.Println(v.ringname)
	}
}