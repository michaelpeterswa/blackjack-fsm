package main

import (
	"context"
	"fmt"

	blackjackfsm "github.com/michaelpeterswa/blackjack-fsm"
	"github.com/michaelpeterswa/blackjack-fsm/deck"
	"github.com/michaelpeterswa/blackjack-fsm/player"
)

func main() {
	blackjackStateMachine := blackjackfsm.NewBlackjackFSM(
		blackjackfsm.WithPlayers(
			player.NewPlayer("Chuck", player.WithInitialBankroll(1000)),
			player.NewPlayer("Sarah", player.WithInitialBankroll(500)),
			player.NewPlayer("Casey", player.WithInitialBankroll(2000)),
			player.NewPlayer("Morgan", player.WithInitialBankroll(100)),
		),
		blackjackfsm.WithDeckOptions(deck.WithNumberOfDecks(6), deck.WithShuffle()),
		blackjackfsm.WithOnEnterStateHandlers(map[blackjackfsm.BlackjackState]func(*blackjackfsm.BlackjackFSM){
			blackjackfsm.StateIdle: func(fsm *blackjackfsm.BlackjackFSM) {
				for _, currentPlayer := range fsm.Players() {
					fmt.Printf("entering idle state for player %s with bankroll: %d\n", currentPlayer.Name, currentPlayer.Bankroll)

					var input int
					fmt.Print("Enter bet: ")
					fmt.Scanln(&input)

					currentPlayer.Bet(input)
					fmt.Printf("Player %s placed a bet of %d. Remaining bankroll: %d\n", currentPlayer.Name, currentPlayer.GetBet(), currentPlayer.Bankroll)
				}
			},
			blackjackfsm.StateDealing: func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("entering dealing") },
			blackjackfsm.StatePlayerTurn: func(fsm *blackjackfsm.BlackjackFSM) {
				for _, currentPlayer := range fsm.Players() {
					fmt.Printf("entering player turn for %s with bet: %d\n", currentPlayer.Name, currentPlayer.GetBet())

					var action string
					fmt.Print("Enter action (hit/stand): ")
					fmt.Scanln(&action)

					if action == "hit" {
						currentPlayer.SetAction(player.ActionHit)
					} else {
						currentPlayer.SetAction(player.ActionStand)
					}
				}
			},
			blackjackfsm.StateDealerTurn: func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("entering dealer turn") },
			blackjackfsm.StateEvaluating: func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("entering evaluating") },
			blackjackfsm.StateGameOver:   func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("entering game over") },
		}),
		blackjackfsm.WithOnExitStateHandlers(map[blackjackfsm.BlackjackState]func(*blackjackfsm.BlackjackFSM){
			blackjackfsm.StateIdle: func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("exiting idle") },
			blackjackfsm.StateDealing: func(fsm *blackjackfsm.BlackjackFSM) {
				for _, player := range fsm.Players() {
					fmt.Printf("exiting dealing for player %s with hand: %v - %d\n", player.Name, player.Hands[0], player.Hands[0].Value())
				}

				fmt.Printf("exiting dealing for dealer with hand: %v - %d\n", fsm.Dealer().Hand(), fsm.Dealer().Hand().Value())
			},
			blackjackfsm.StatePlayerTurn: func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("exiting player turn") },
			blackjackfsm.StateDealerTurn: func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("exiting dealer turn") },
			blackjackfsm.StateEvaluating: func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("exiting evaluating") },
			blackjackfsm.StateGameOver:   func(fsm *blackjackfsm.BlackjackFSM) { fmt.Println("exiting game over") },
		}),
	)

	fmt.Println("Blackjack FSM initialized with state:", blackjackStateMachine)

	blackjackStateMachine.Run(context.Background())
}
