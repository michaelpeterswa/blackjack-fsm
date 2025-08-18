package blackjackfsm

import (
	"context"
	"fmt"
	"time"

	"github.com/michaelpeterswa/blackjack-fsm/dealer"
	"github.com/michaelpeterswa/blackjack-fsm/deck"
	"github.com/michaelpeterswa/blackjack-fsm/player"
)

type BlackjackFSM struct {
	currentState  BlackjackState
	deckOptions   []deck.DeckOption
	currentDeck   *deck.Deck
	activePlayers []*player.Player
	dealer        *dealer.Dealer

	onEnterStateHandlers map[BlackjackState]func(*BlackjackFSM)
	onExitStateHandlers  map[BlackjackState]func(*BlackjackFSM)
}

type BlackjackFSMOption func(*BlackjackFSM)

func NewBlackjackFSM(options ...BlackjackFSMOption) *BlackjackFSM {
	fsm := &BlackjackFSM{
		currentState: StateIdle,
		deckOptions:  []deck.DeckOption{},
		dealer:       dealer.NewDealer(),
	}
	for _, option := range options {
		option(fsm)
	}

	fsm.currentDeck = deck.NewDeck(fsm.deckOptions...)
	if fsm.currentDeck == nil {
		fsm.currentDeck = deck.NewDeck(deck.WithNumberOfDecks(1), deck.WithShuffle())
	}

	return fsm
}

func WithDeckOptions(options ...deck.DeckOption) BlackjackFSMOption {
	return func(fsm *BlackjackFSM) {
		fsm.deckOptions = options
	}
}

func WithPlayers(players ...*player.Player) BlackjackFSMOption {
	return func(fsm *BlackjackFSM) {
		fsm.activePlayers = players
	}
}

func WithOnEnterStateHandlers(handlers map[BlackjackState]func(*BlackjackFSM)) BlackjackFSMOption {
	return func(fsm *BlackjackFSM) {
		fsm.onEnterStateHandlers = handlers
	}
}

func WithOnExitStateHandlers(handlers map[BlackjackState]func(*BlackjackFSM)) BlackjackFSMOption {
	return func(fsm *BlackjackFSM) {
		fsm.onExitStateHandlers = handlers
	}
}

func (fsm *BlackjackFSM) SetState(state BlackjackState) {
	fsm.currentState = state
}

func (fsm *BlackjackFSM) State() BlackjackState {
	return fsm.currentState
}

func (fsm *BlackjackFSM) Players() []*player.Player {
	return fsm.activePlayers
}

func (fsm *BlackjackFSM) Dealer() *dealer.Dealer {
	return fsm.dealer
}

func (fsm *BlackjackFSM) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := fsm.handleState(); err != nil {
				return err
			}
		}
		time.Sleep(1 * time.Second) // Simulate time passing for state transitions
	}
}

func (fsm *BlackjackFSM) handleState() error {
	// on enter callback
	enteredState := fsm.currentState
	if handler, exists := fsm.onEnterStateHandlers[enteredState]; exists {
		handler(fsm)
	}

	switch fsm.currentState {
	case StateIdle:
		fsm.Idle()
		// Transition to dealing state
		fsm.SetState(StateDealing)
	case StateDealing:
		fsm.Dealing()
		// Logic to deal cards
		// Transition to player's turn state
		fsm.SetState(StatePlayerTurn)
	case StatePlayerTurn:
		fsm.PlayerTurn()
		// Logic for player's actions (hit, stand, etc.)
		// After player's turn, transition to dealer's turn state
		fsm.SetState(StateDealerTurn)
	case StateDealerTurn:
		fsm.DealerTurn()
		// Logic for dealer's actions
		// After dealer's turn, transition to evaluating state
		fsm.SetState(StateEvaluating)
	case StateEvaluating:
		fsm.Evaluating()

		fsm.SetState(StateGameOver)
	case StateGameOver:
		fsm.GameOver()
		// Logic to reset the game
		fsm.SetState(StateIdle)
	default:
		return fmt.Errorf("unknown state: %v", fsm.currentState)
	}

	// ending callback
	if handler, exists := fsm.onExitStateHandlers[enteredState]; exists {
		handler(fsm)
	}

	return nil
}
