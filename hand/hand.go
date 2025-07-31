package hand

import "github.com/michaelpeterswa/blackjack-fsm/card"

type Hand []card.Card

func (h *Hand) AddCard(card card.Card) {
	*h = append(*h, card)
}

func (h *Hand) Clear() {
	*h = []card.Card{}
}

func (h *Hand) Aces() int {
	count := 0
	for _, c := range *h {
		if c.Rank == card.RankAce {
			count++
		}
	}
	return count
}

func (h *Hand) Value() []int {
	var values []int
	for _, c := range *h {
		if len(values) == 0 { // no values exist yet
			values = append(values, c.Rank.Values()...)
			continue
		}
		if c.Rank == card.RankAce {
			if len(values) == 1 && h.Aces() == 1 { // need to add high ace as discrete value
				valuePreAdd := values[0]
				values[0] += c.Rank.Values()[0]
				if valuePreAdd < 11 { // if the previous value was less than 11, we can add the low ace
					values = append(values, c.Rank.Values()[1]+valuePreAdd)
				}
			} else if len(values) == 2 { // already have two values, just add the low ace
				for i := range values {
					values[i] += c.Rank.Values()[0]
				}
			}
		} else {

			for i := range values {
				if values[i]+c.Rank.Values()[0] > 21 && len(values) != 1 {
					// remove this value if it exceeds 21
					values = append(values[:i], values[i+1:]...)
					// i-- was here but removed for ineffassign
					continue
				}
				values[i] += c.Rank.Values()[0]
			}
		}
	}
	return values
}
