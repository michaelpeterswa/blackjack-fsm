package blackjackfsm

type Hand []Card

type Player struct {
	Name     string
	Hands    []Hand
	Bankroll int
}

type PlayerOption func(*Player) *Player

func NewPlayer(name string, options ...PlayerOption) *Player {
	player := &Player{
		Name:     name,
		Hands:    []Hand{},
		Bankroll: 1000, // Default bankroll
	}

	for _, opt := range options {
		player = opt(player)
	}

	return player
}

func WithInitialBankroll(bankroll int) PlayerOption {
	return func(p *Player) *Player {
		p.Bankroll = bankroll
		return p
	}
}

func (h *Hand) AddCard(card Card) {
	*h = append(*h, card)
}

func (p *Player) ClearHands() {
	p.Hands = []Hand{}
}
