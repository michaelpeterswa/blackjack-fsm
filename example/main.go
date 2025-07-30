package main

import (
	"fmt"

	blackjackfsm "github.com/michaelpeterswa/blackjack-fsm"
)

func main() {
	blackjackstatemachine := blackjackfsm.NewBlackjackFSM()

	fmt.Println("Blackjack FSM initialized with state:", blackjackstatemachine)
}
