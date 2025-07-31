package player

import (
	"github.com/michaelpeterswa/blackjack-fsm/hand"
)

type Player struct {
	Name          string
	Hands         []hand.Hand
	Bankroll      int
	currentBet    int
	currentAction PlayerAction
}

type PlayerOption func(*Player) *Player

func NewPlayer(name string, options ...PlayerOption) *Player {
	player := &Player{
		Name:     name,
		Hands:    []hand.Hand{},
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

func (p *Player) ClearHands() {
	p.Hands = []hand.Hand{}
}

func (p *Player) Bet(amount int) {
	if amount > p.Bankroll {
		amount = p.Bankroll // Cannot bet more than bankroll
	}
	p.currentBet = amount
	p.Bankroll -= amount
}

func (p *Player) GetBet() int {
	return p.currentBet
}

func (p *Player) Credit(amount int) {
	p.Bankroll += amount
}

type PlayerAction int

const (
	ActionHit PlayerAction = iota
	ActionStand
	ActionDoubleDown
	ActionSplit
	ActionInsurance
)

func (p *Player) SetAction(action PlayerAction) {
	p.currentAction = action
}

func (p *Player) Action() PlayerAction {
	return p.currentAction
}
