package dealer

import (
	"github.com/michaelpeterswa/blackjack-fsm/card"
	"github.com/michaelpeterswa/blackjack-fsm/hand"
)

type Dealer struct {
	hand          *hand.Hand
	mustHitSoft17 bool
}

type DealerOption func(*Dealer)

func NewDealer(options ...DealerOption) *Dealer {
	dealer := &Dealer{
		hand:          &hand.Hand{},
		mustHitSoft17: true, // dealer hits on soft 17 by default
	}
	for _, opt := range options {
		opt(dealer)
	}
	return dealer
}

func WithMustHitSoft17(mustHit bool) DealerOption {
	return func(d *Dealer) {
		d.mustHitSoft17 = mustHit
	}
}

func (d *Dealer) AddCardToHand(card card.Card) {
	d.hand.AddCard(card)
}

func (d *Dealer) ClearHand() {
	d.hand.Clear()
}

func (d *Dealer) Hand() *hand.Hand {
	return d.hand
}

// TODO: not provably correct yet, needs more testing
func (d *Dealer) MustHit() bool {
	// Dealer must hit if hand value is less than 17 or if it has a soft 17 (Ace + 6)
	return d.hand.Value()[0] < 17 || ((d.hand.Aces() > 0 && d.hand.Value()[0] == 17) && d.mustHitSoft17)
}
