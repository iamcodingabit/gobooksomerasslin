package roster

import (
	
)

func addWrestler(wrestler Wrestler, ringname, alignment, signature_move string) Wrestler{
	wrestler.ringname = ringname
	wrestler.alignment = alignment
	wrestler.signature_move = signature_move

	return wrestler
}

func listWrestler(ringname string) map[int]Wrestler{
	//Some get statement here
	//Map to Wrestler

	wrestlers := map[int]Wrestler{

	}

	return wrestlers
}

func removeWrestler(wrestler Wrestler){
	
}

func editWrestler(wrestler Wrestler){
	
}