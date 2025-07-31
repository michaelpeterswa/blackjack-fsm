package hand_test

import (
	"testing"

	"github.com/michaelpeterswa/blackjack-fsm/card"
	"github.com/michaelpeterswa/blackjack-fsm/hand"
)

func TestValue(t *testing.T) {
	tests := []struct {
		name           string
		cards          []card.Card
		expectedValues []int
	}{
		// {
		// 	name: "Ace and Ten",
		// 	cards: []card.Card{
		// 		{Rank: card.RankAce, Suit: card.SuitHearts},
		// 		{Rank: card.RankTen, Suit: card.SuitDiamonds},
		// 	},
		// 	expectedValues: []int{11, 21},
		// },
		// {
		// 	name: "Ten and Ace",
		// 	cards: []card.Card{
		// 		{Rank: card.RankTen, Suit: card.SuitDiamonds},
		// 		{Rank: card.RankAce, Suit: card.SuitHearts},
		// 	},
		// 	expectedValues: []int{11, 21},
		// },
		// {
		// 	name: "Two Aces",
		// 	cards: []card.Card{
		// 		{Rank: card.RankAce, Suit: card.SuitHearts},
		// 		{Rank: card.RankAce, Suit: card.SuitDiamonds},
		// 	},
		// 	expectedValues: []int{2, 12},
		// },
		// {
		// 	name: "Ten and King",
		// 	cards: []card.Card{
		// 		{Rank: card.RankTen, Suit: card.SuitHearts},
		// 		{Rank: card.RankKing, Suit: card.SuitDiamonds},
		// 	},
		// 	expectedValues: []int{20},
		// },
		// {
		// 	name: "Two and Three",
		// 	cards: []card.Card{
		// 		{Rank: card.RankTwo, Suit: card.SuitHearts},
		// 		{Rank: card.RankThree, Suit: card.SuitDiamonds},
		// 	},
		// 	expectedValues: []int{5},
		// },
		// {
		// 	name: "Single Ace",
		// 	cards: []card.Card{
		// 		{Rank: card.RankAce, Suit: card.SuitHearts},
		// 	},
		// 	expectedValues: []int{1, 11},
		// },
		// {
		// 	name: "Ace Hits",
		// 	cards: []card.Card{
		// 		{Rank: card.RankAce, Suit: card.SuitHearts},
		// 		{Rank: card.RankTwo, Suit: card.SuitDiamonds},
		// 		{Rank: card.RankKing, Suit: card.SuitClubs},
		// 	},
		// 	expectedValues: []int{13},
		// },
		{
			name: "Bust",
			cards: []card.Card{
				{Rank: card.RankTen, Suit: card.SuitHearts},
				{Rank: card.RankJack, Suit: card.SuitDiamonds},
				{Rank: card.RankQueen, Suit: card.SuitClubs},
			},
			expectedValues: []int{30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand := &hand.Hand{}
			for _, card := range tt.cards {
				hand.AddCard(card)
			}

			values := hand.Value()

			if len(values) != len(tt.expectedValues) {
				t.Errorf("Expected %d values, got %d", len(tt.expectedValues), len(values))
			}

			for i, v := range values {
				if v != tt.expectedValues[i] {
					t.Errorf("Expected value %d at index %d, got %d", tt.expectedValues[i], i, v)
				}
			}
		})
	}
}
