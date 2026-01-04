package main

import (
	"bufio"
	"fmt"
	"os"
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
	
	fmt.Print("Keep adding by typing 'add': ")
	input := ""
	isActive := true
	//Adding wrestlers
	

	for isActive {
		scanner.Scan()
		input = scanner.Text()
		if input == "add" {
			Commands[input].(func(*[]Wrestler))(&wrestlers)
			fmt.Println(wrestlers)
		} else {
			isActive = false
		}
	}
	
	fmt.Println("Thanks for using the app!")
}

func addWrestler(wrestlers *[]Wrestler){
	var var1 Wrestler
	var1.ringname = "Billy Graham"
	var1.nickname = "Superstar"
	var1.signature = "Suplex"
	var1.finisher = "Powerslam"
	
	*wrestlers = append(*wrestlers, var1)
}