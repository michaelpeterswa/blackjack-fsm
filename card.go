package blackjackfsm

import (
	"fmt"
	"math/rand/v2"
)

type Suit int

type Rank int

const (
	SuitHearts Suit = iota
	SuitDiamonds
	SuitClubs
	SuitSpades
)

func (s Suit) String() string {
	switch s {
	case SuitHearts:
		return "Hearts"
	case SuitDiamonds:
		return "Diamonds"
	case SuitClubs:
		return "Clubs"
	case SuitSpades:
		return "Spades"
	default:
		return "Unknown"
	}
}

const (
	RankTwo Rank = iota + 2
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
)

func (r Rank) String() string {
	switch r {
	case RankTwo:
		return "2"
	case RankThree:
		return "3"
	case RankFour:
		return "4"
	case RankFive:
		return "5"
	case RankSix:
		return "6"
	case RankSeven:
		return "7"
	case RankEight:
		return "8"
	case RankNine:
		return "9"
	case RankTen:
		return "10"
	case RankJack:
		return "J"
	case RankQueen:
		return "Q"
	case RankKing:
		return "K"
	case RankAce:
		return "A"
	default:
		return "Unknown"
	}
}

func (r Rank) Values() []int {
	switch r {
	case RankTwo:
		return []int{2}
	case RankThree:
		return []int{3}
	case RankFour:
		return []int{4}
	case RankFive:
		return []int{5}
	case RankSix:
		return []int{6}
	case RankSeven:
		return []int{7}
	case RankEight:
		return []int{8}
	case RankNine:
		return []int{9}
	case RankTen, RankJack, RankQueen, RankKing:
		return []int{10}
	case RankAce:
		return []int{1, 11} // Ace can be 1 or 11
	default:
		return nil
	}
}

type Card struct {
	Suit Suit
	Rank Rank
}

var (
	// Hearts
	CardTwoHearts   = Card{Suit: SuitHearts, Rank: RankTwo}
	CardThreeHearts = Card{Suit: SuitHearts, Rank: RankThree}
	CardFourHearts  = Card{Suit: SuitHearts, Rank: RankFour}
	CardFiveHearts  = Card{Suit: SuitHearts, Rank: RankFive}
	CardSixHearts   = Card{Suit: SuitHearts, Rank: RankSix}
	CardSevenHearts = Card{Suit: SuitHearts, Rank: RankSeven}
	CardEightHearts = Card{Suit: SuitHearts, Rank: RankEight}
	CardNineHearts  = Card{Suit: SuitHearts, Rank: RankNine}
	CardTenHearts   = Card{Suit: SuitHearts, Rank: RankTen}
	CardJackHearts  = Card{Suit: SuitHearts, Rank: RankJack}
	CardQueenHearts = Card{Suit: SuitHearts, Rank: RankQueen}
	CardKingHearts  = Card{Suit: SuitHearts, Rank: RankKing}
	CardAceHearts   = Card{Suit: SuitHearts, Rank: RankAce}

	// Diamonds
	CardTwoDiamonds   = Card{Suit: SuitDiamonds, Rank: RankTwo}
	CardThreeDiamonds = Card{Suit: SuitDiamonds, Rank: RankThree}
	CardFourDiamonds  = Card{Suit: SuitDiamonds, Rank: RankFour}
	CardFiveDiamonds  = Card{Suit: SuitDiamonds, Rank: RankFive}
	CardSixDiamonds   = Card{Suit: SuitDiamonds, Rank: RankSix}
	CardSevenDiamonds = Card{Suit: SuitDiamonds, Rank: RankSeven}
	CardEightDiamonds = Card{Suit: SuitDiamonds, Rank: RankEight}
	CardNineDiamonds  = Card{Suit: SuitDiamonds, Rank: RankNine}
	CardTenDiamonds   = Card{Suit: SuitDiamonds, Rank: RankTen}
	CardJackDiamonds  = Card{Suit: SuitDiamonds, Rank: RankJack}
	CardQueenDiamonds = Card{Suit: SuitDiamonds, Rank: RankQueen}
	CardKingDiamonds  = Card{Suit: SuitDiamonds, Rank: RankKing}
	CardAceDiamonds   = Card{Suit: SuitDiamonds, Rank: RankAce}

	// Clubs
	CardTwoClubs   = Card{Suit: SuitClubs, Rank: RankTwo}
	CardThreeClubs = Card{Suit: SuitClubs, Rank: RankThree}
	CardFourClubs  = Card{Suit: SuitClubs, Rank: RankFour}
	CardFiveClubs  = Card{Suit: SuitClubs, Rank: RankFive}
	CardSixClubs   = Card{Suit: SuitClubs, Rank: RankSix}
	CardSevenClubs = Card{Suit: SuitClubs, Rank: RankSeven}
	CardEightClubs = Card{Suit: SuitClubs, Rank: RankEight}
	CardNineClubs  = Card{Suit: SuitClubs, Rank: RankNine}
	CardTenClubs   = Card{Suit: SuitClubs, Rank: RankTen}
	CardJackClubs  = Card{Suit: SuitClubs, Rank: RankJack}
	CardQueenClubs = Card{Suit: SuitClubs, Rank: RankQueen}
	CardKingClubs  = Card{Suit: SuitClubs, Rank: RankKing}
	CardAceClubs   = Card{Suit: SuitClubs, Rank: RankAce}

	// Spades
	CardTwoSpades   = Card{Suit: SuitSpades, Rank: RankTwo}
	CardThreeSpades = Card{Suit: SuitSpades, Rank: RankThree}
	CardFourSpades  = Card{Suit: SuitSpades, Rank: RankFour}
	CardFiveSpades  = Card{Suit: SuitSpades, Rank: RankFive}
	CardSixSpades   = Card{Suit: SuitSpades, Rank: RankSix}
	CardSevenSpades = Card{Suit: SuitSpades, Rank: RankSeven}
	CardEightSpades = Card{Suit: SuitSpades, Rank: RankEight}
	CardNineSpades  = Card{Suit: SuitSpades, Rank: RankNine}
	CardTenSpades   = Card{Suit: SuitSpades, Rank: RankTen}
	CardJackSpades  = Card{Suit: SuitSpades, Rank: RankJack}
	CardQueenSpades = Card{Suit: SuitSpades, Rank: RankQueen}
	CardKingSpades  = Card{Suit: SuitSpades, Rank: RankKing}
	CardAceSpades   = Card{Suit: SuitSpades, Rank: RankAce}
)

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

type Deck struct {
	cards     []Card
	deckCount int
	doShuffle bool
}

type DeckOption func(d Deck) Deck

var (
	StandardSingleDeck = []Card{
		// Hearts
		CardTwoHearts, CardThreeHearts, CardFourHearts, CardFiveHearts, CardSixHearts,
		CardSevenHearts, CardEightHearts, CardNineHearts, CardTenHearts, CardJackHearts,
		CardQueenHearts, CardKingHearts, CardAceHearts,

		// Diamonds
		CardTwoDiamonds, CardThreeDiamonds, CardFourDiamonds, CardFiveDiamonds, CardSixDiamonds,
		CardSevenDiamonds, CardEightDiamonds, CardNineDiamonds, CardTenDiamonds, CardJackDiamonds,
		CardQueenDiamonds, CardKingDiamonds, CardAceDiamonds,

		// Clubs
		CardTwoClubs, CardThreeClubs, CardFourClubs, CardFiveClubs, CardSixClubs,
		CardSevenClubs, CardEightClubs, CardNineClubs, CardTenClubs, CardJackClubs,
		CardQueenClubs, CardKingClubs, CardAceClubs,

		// Spades
		CardTwoSpades, CardThreeSpades, CardFourSpades, CardFiveSpades, CardSixSpades,
		CardSevenSpades, CardEightSpades, CardNineSpades, CardTenSpades, CardJackSpades,
		CardQueenSpades, CardKingSpades, CardAceSpades,
	}
)

func WithNumberOfDecks(n int) DeckOption {
	return func(d Deck) Deck {
		d.deckCount = n
		return d
	}
}

func WithShuffle() DeckOption {
	return func(d Deck) Deck {
		d.doShuffle = true
		return d
	}
}

func NewDeck(opts ...DeckOption) *Deck {
	var d Deck
	defaultDeckCount := 1

	if d.deckCount == 0 {
		d.deckCount = defaultDeckCount
	}

	for _, opt := range opts {
		d = opt(d)
	}

	d.cards = make([]Card, 0, len(StandardSingleDeck)*d.deckCount)
	for i := 0; i < d.deckCount; i++ {
		d.cards = append(d.cards, StandardSingleDeck...)
	}

	if d.doShuffle {
		d.Shuffle()
	}

	return &d
}

func (d *Deck) Len() int {
	return len(d.cards)
}

func (d *Deck) Draw() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, fmt.Errorf("deck is empty")
	}

	card := d.cards[0]
	d.cards = d.cards[1:]

	return card, nil
}

func (d *Deck) Cards() []Card {
	return d.cards
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Cut(n int) error {
	if n < 1 || n > len(d.cards) {
		return fmt.Errorf("cut size %d is out of bounds for deck of size %d", n, len(d.cards))
	}

	tmpDeck := make([]Card, 0, len(d.cards))

	// Move the first n cards to the end of the deck
	tmpDeck = append(tmpDeck, d.cards[n:]...)
	tmpDeck = append(tmpDeck, d.cards[:n]...)

	d.cards = tmpDeck

	return nil
}
