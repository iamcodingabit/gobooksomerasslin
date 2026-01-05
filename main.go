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
}
	


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	wrestlers := []Wrestler{}
	
	fmt.Print("Add wrestlers by typing 'add', type 'quit' to quit: ")
	input := ""
	isActive := true
	//Adding wrestlers
	

	for isActive {
		scanner.Scan()
		input = scanner.Text()
		command := strings.SplitAfterN(input, " ", 2)
		fmt.Printf("\"%s\": %T\n", command[0], command[0])
		fmt.Printf("\"%s\": %T\n", command[1], command[1])
		if input != "quit" {
			Commands[strings.Trim(command[0], " ")].(func(*[]Wrestler, string))(&wrestlers, command[1])
			fmt.Println(wrestlers)
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