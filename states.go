package blackjackfsm

import (
	"fmt"
	"math/rand"

	"github.com/michaelpeterswa/blackjack-fsm/deck"
	"github.com/michaelpeterswa/blackjack-fsm/hand"
	"github.com/michaelpeterswa/blackjack-fsm/player"
)

type BlackjackState int

const (
	StateIdle BlackjackState = iota
	StateDealing
	StatePlayerTurn
	StateDealerTurn
	StateEvaluating
	StateGameOver
)

func (fsm *BlackjackFSM) Idle() {
	fmt.Println("in idle state")

	for _, player := range fsm.activePlayers {
		player.ClearHands()
	}

	fsm.currentDeck = deck.NewDeck(fsm.deckOptions...)
	fsm.currentDeck.Cut(rand.Intn(fsm.currentDeck.Len()))

	// take bets from players
}

func (fsm *BlackjackFSM) Dealing() {
	fmt.Println("in dealing state")

	// PLAYER
	for _, player := range fsm.activePlayers {
		if len(player.Hands) == 0 {
			player.Hands = append(player.Hands, hand.Hand{})
		}

		for i := 0; i < 2; i++ {
			card, err := fsm.currentDeck.Draw()
			if err != nil {
				fmt.Println("No more cards in the deck to deal.")
				return
			}
			player.Hands[0].AddCard(card)
		}
	}

	// DEALER
	fsm.dealer.ClearHand()
	for i := 0; i < 2; i++ {
		card, err := fsm.currentDeck.Draw()
		if err != nil {
			fmt.Println("No more cards in the deck to deal.")
			return
		}
		fsm.dealer.AddCardToHand(card)
	}
}

func (fsm *BlackjackFSM) PlayerTurn() {
	fmt.Println("in player turn state")

	for _, currentPlayer := range fsm.activePlayers {
		if len(currentPlayer.Hands) == 0 {
			continue // Skip players with no hands
		}

		switch currentPlayer.Action() {
		case player.ActionHit:
			card, err := fsm.currentDeck.Draw()
			if err != nil {
				fmt.Println("No more cards in the deck to hit.")
				return
			}

			currentPlayer.Hands[0].AddCard(card)
			fmt.Printf("Player %s hits and receives %s. Hand value: %d\n", currentPlayer.Name, card, currentPlayer.Hands[0].Value())
		case player.ActionStand:
			fmt.Printf("Player %s stands with hand value: %d\n", currentPlayer.Name, currentPlayer.Hands[0].Value())
		default:
			fmt.Printf("Player %s has no action set.\n", currentPlayer.Name)
		}
	}
}

func (fsm *BlackjackFSM) DealerTurn() {
	fmt.Println("in dealer turn state")

	for fsm.dealer.MustHit() {
		card, err := fsm.currentDeck.Draw()
		if err != nil {
			fmt.Println("No more cards in the deck for dealer to hit.")
			return
		}
		fsm.dealer.AddCardToHand(card)
		fmt.Printf("Dealer hits and receives %s. Hand value: %d\n", card, fsm.dealer.Hand().Value())
	}

	fmt.Printf("Dealer's hand: %v = %d\n", fsm.dealer.Hand(), fsm.dealer.Hand().Value())
}

func (fsm *BlackjackFSM) Evaluating() {
	fmt.Println("in evaluate state")

	for _, currentPlayer := range fsm.activePlayers {
		if len(currentPlayer.Hands) == 0 {
			continue // Skip players with no hands
		}

		playerValue := currentPlayer.Hands[0].Value()[0]
		dealerValue := fsm.dealer.Hand().Value()[0]

		if playerValue > 21 {
			fmt.Printf("Player %s busts with hand value: %d\n", currentPlayer.Name, playerValue)
		} else if dealerValue > 21 || playerValue > dealerValue {
			fmt.Printf("Player %s wins with hand value: %d against dealer's %d\n", currentPlayer.Name, playerValue, dealerValue)
			currentPlayer.Credit(currentPlayer.GetBet() * 2) // Player wins and gets double their bet
		} else if playerValue < dealerValue {
			fmt.Printf("Dealer wins against player %s with hand value: %d against player's %d\n", currentPlayer.Name, dealerValue, playerValue)
		} else {
			fmt.Printf("It's a tie for player %s with hand value: %d\n", currentPlayer.Name, playerValue)
			currentPlayer.Credit(currentPlayer.GetBet()) // Player gets their bet back
		}
	}
}

func (fsm *BlackjackFSM) GameOver() {
	fmt.Println("in game over state")
}
