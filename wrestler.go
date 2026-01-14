package wrestler

import (
	"fmt"
)

type Wrestler struct {
	ringname string
	nickname string
	alignment string
	signature string
	finisher string
}

func addWrestler(wrestlers *[]Wrestler, newWrestler string) string{
	return newWrestler
}

func listWrestlers(wrestlers []Wrestler){
	for _, v := range wrestlers {
		fmt.Println(v.ringname)
	}
}