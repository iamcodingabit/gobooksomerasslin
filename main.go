package main

import (
	"fmt"
)

func main(){
	wrestlers = append(wrestlers, 
		Wrestler{ringname: "Kurt Angle", alignment: "heel", signature_move: "Ankle Lock"},
		Wrestler{ringname: "Hiroshi Tanahashi", alignment: "face", signature_move: "High Fly Flow"},
		Wrestler{ringname: "Bryan Danielson", alignment: "heel", signature_move: "Lebell Lock"})
	
	for i := 0; i < len(wrestlers); i++ {
		if wrestlers[i].ringname == "Kurt Angle"{
			wrestlers[i].ringname = "Chad Gable"
			break
		}
	}
	
	fmt.Println("Ringname\tAlignment\tSignature Move")
	for i := 0; i < len(wrestlers); i++{
		fmt.Printf("%s\t%s\t%s\n", 
			wrestlers[i].ringname,
			wrestlers[i].alignment,
			wrestlers[i].signature_move,
		)
	}
	
}