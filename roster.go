package main

import "fmt"

func addWrestler(wrestler Wrestler){
	wrestlers = append(wrestlers, wrestler)
}

func listWrestlers(){
	for _, y := range wrestlers {
		fmt.Printf("%s - %s - %s\n", y.ringname, y.alignment, y.signature_move)
	}
}

func removeWrestler(wrestler Wrestler){
	
}

func editWrestler(wrestler Wrestler){
	
}