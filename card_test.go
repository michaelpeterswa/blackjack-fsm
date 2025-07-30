package blackjackfsm_test

import (
	"testing"

	blackjackfsm "github.com/michaelpeterswa/blackjack-fsm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDeck(t *testing.T) {
	tests := []struct {
		name        string
		numDecks    int
		expectedLen int
	}{
		{"one deck", 1, 52},
		{"two decks", 2, 104},
		{"six decks", 6, 312},
		{"zero decks", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := blackjackfsm.NewDeck(blackjackfsm.WithNumberOfDecks(tt.numDecks))
			assert.NotNil(t, deck)
			assert.Equal(t, tt.expectedLen, deck.Len())
		})
	}
}

func TestDeckShuffle(t *testing.T) {
	deck := blackjackfsm.NewDeck(blackjackfsm.WithNumberOfDecks(1), blackjackfsm.WithShuffle())
	nonShuffledDeck := blackjackfsm.NewDeck(blackjackfsm.WithNumberOfDecks(1))

	require.NotNil(t, deck)
	require.NotNil(t, nonShuffledDeck)

	// ensure the deck has 52 cards
	assert.Equal(t, 52, deck.Len())
	assert.Equal(t, 52, nonShuffledDeck.Len())

	// ensure order is different after shuffle
	assert.NotEqual(t, nonShuffledDeck.Cards(), deck.Cards())
	assert.ElementsMatch(t, nonShuffledDeck.Cards(), deck.Cards())
}
