package blackjackfsm

type BlackjackState int

const (
	StateIdle BlackjackState = iota
	StateDealing
	StatePlayerTurn
	StateDealerTurn
	StateGameOver
)

type PlayerAction int

const (
	ActionHit PlayerAction = iota
	ActionStand
	ActionDoubleDown
	ActionSplit
	ActionInsurance
)

type BlackjackFSM struct {
	currentState BlackjackState
	currentDeck  *Deck
}

func NewBlackjackFSM() *BlackjackFSM {
	return &BlackjackFSM{
		currentState: StateIdle,
		currentDeck:  NewDeck(WithNumberOfDecks(1), WithShuffle()),
	}
}
