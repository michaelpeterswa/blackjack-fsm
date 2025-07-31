package card_test

import (
	"testing"

	"github.com/michaelpeterswa/blackjack-fsm/card"
)

func TestCardString(t *testing.T) {
	card := card.Card{Rank: card.RankAce, Suit: card.SuitHearts}
	expected := "Ace of Hearts"
	if card.String() != expected {
		t.Errorf("Expected %s, got %s", expected, card.String())
	}
}
