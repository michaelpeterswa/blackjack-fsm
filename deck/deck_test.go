package deck_test

import (
	"testing"

	"github.com/michaelpeterswa/blackjack-fsm/card"
	"github.com/michaelpeterswa/blackjack-fsm/deck"
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
			deck := deck.NewDeck(deck.WithNumberOfDecks(tt.numDecks))
			assert.NotNil(t, deck)
			assert.Equal(t, tt.expectedLen, deck.Len())
		})
	}
}

func TestDeckShuffle(t *testing.T) {
	shuffledDeck := deck.NewDeck(deck.WithNumberOfDecks(1), deck.WithShuffle())
	nonShuffledDeck := deck.NewDeck(deck.WithNumberOfDecks(1))

	require.NotNil(t, shuffledDeck)
	require.NotNil(t, nonShuffledDeck)

	// ensure the deck has 52 cards
	assert.Equal(t, 52, shuffledDeck.Len())
	assert.Equal(t, 52, nonShuffledDeck.Len())

	// ensure order is different after shuffle
	assert.NotEqual(t, nonShuffledDeck.Cards(), shuffledDeck.Cards())
	assert.ElementsMatch(t, nonShuffledDeck.Cards(), shuffledDeck.Cards())
}

func TestDeckCut(t *testing.T) {
	preCutDeck := deck.NewDeck(deck.WithNumberOfDecks(1))

	tests := []struct {
		name      string
		cutSize   int
		expectErr bool
	}{
		{"valid cut", 10, false},
		{"cut size too small", 0, true},
		{"cut size too large", 53, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := preCutDeck.Cut(tt.cutSize)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeckDraw(t *testing.T) {
	deck := deck.NewDeck(deck.WithNumberOfDecks(1))

	// Draw all cards from the deck
	for i := 0; i < 52; i++ {
		drawCard, err := deck.Draw()
		assert.NoError(t, err)
		assert.NotEqual(t, card.Card{}, drawCard)
	}

	// Attempt to draw from an empty deck
	_, err := deck.Draw()
	assert.Error(t, err)
	assert.Equal(t, "deck is empty", err.Error())
}
