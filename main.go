package main

import (
	//"bufio"
	//"os"
	//"strings"
	"context"
	"fmt"
	"github.com/iamcodingabit/gobooksomerasslin/database"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	pool := database.connect()
	c := context.Background()

	w := wrestlerRepository {
		dbpool: pool,		
	}
	wrestlerToUpdate := Wrestler{"Kurt Angle", "face", "Ankle Lock"} //Whoops! I meant to do Chad Gable
	newWrestler := "Chad Gable"
	wrestlerToDelete := Wrestler{"Contra Verci", "heel", "Paws of Injustice"} //Can't keep this guy around...

	w.Create(c, &wrestlerToUpdate)
	w.Create(c, &wrestlerToDelete)
	fmt.Println()
	fmt.Println("Wrestlers after creations (to update and to delete):")
	w.ReadAllWrestler()
	
	fmt.Println()
	fmt.Println(w.DeleteByRingname(c, wrestlerToDelete.Ringname))
	fmt.Println(w.UpdateRingname(c, "Kurt Angle", "Chad Gable"))
	fmt.Println()
	
	fmt.Println("After updating and deleting:")
	w.ReadAllWrestler()
	fmt.Println()
	//closing, deleting the changes to just the seeded entries to repeat
	w.DeleteByRingname(c, newWrestler)
	defer pool.Close()
}