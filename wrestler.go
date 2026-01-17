package roster

import (
	//"fmt"
)

type Wrestler struct {
	ringname string
	alignment string // face | heel | tweener?
	signature_move string
}

type WrestlerBehavior interface {
	
}