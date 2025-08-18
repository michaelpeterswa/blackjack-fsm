package deck

import (
	"fmt"
	"math/rand"

	"github.com/michaelpeterswa/blackjack-fsm/card"
)

type Deck struct {
	cards     []card.Card
	deckCount int
	doShuffle bool
}

type DeckOption func(d Deck) Deck

var (
	StandardSingleDeck = []card.Card{
		// Hearts
		card.CardTwoHearts, card.CardThreeHearts, card.CardFourHearts, card.CardFiveHearts, card.CardSixHearts,
		card.CardSevenHearts, card.CardEightHearts, card.CardNineHearts, card.CardTenHearts, card.CardJackHearts,
		card.CardQueenHearts, card.CardKingHearts, card.CardAceHearts,

		// Diamonds
		card.CardTwoDiamonds, card.CardThreeDiamonds, card.CardFourDiamonds, card.CardFiveDiamonds, card.CardSixDiamonds,
		card.CardSevenDiamonds, card.CardEightDiamonds, card.CardNineDiamonds, card.CardTenDiamonds, card.CardJackDiamonds,
		card.CardQueenDiamonds, card.CardKingDiamonds, card.CardAceDiamonds,

		// Clubs
		card.CardTwoClubs, card.CardThreeClubs, card.CardFourClubs, card.CardFiveClubs, card.CardSixClubs,
		card.CardSevenClubs, card.CardEightClubs, card.CardNineClubs, card.CardTenClubs, card.CardJackClubs,
		card.CardQueenClubs, card.CardKingClubs, card.CardAceClubs,

		// Spades
		card.CardTwoSpades, card.CardThreeSpades, card.CardFourSpades, card.CardFiveSpades, card.CardSixSpades,
		card.CardSevenSpades, card.CardEightSpades, card.CardNineSpades, card.CardTenSpades, card.CardJackSpades,
		card.CardQueenSpades, card.CardKingSpades, card.CardAceSpades,
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

	d.cards = make([]card.Card, 0, len(StandardSingleDeck)*d.deckCount)
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

func (d *Deck) Draw() (card.Card, error) {
	if len(d.cards) == 0 {
		return card.Card{}, fmt.Errorf("deck is empty")
	}

	card := d.cards[0]
	d.cards = d.cards[1:]

	return card, nil
}

func (d *Deck) Cards() []card.Card {
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

	tmpDeck := make([]card.Card, 0, len(d.cards))

	// Move the first n cards to the end of the deck
	tmpDeck = append(tmpDeck, d.cards[n:]...)
	tmpDeck = append(tmpDeck, d.cards[:n]...)

	d.cards = tmpDeck

	return nil
}
