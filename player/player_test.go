package player_test

import (
	"testing"

	"github.com/michaelpeterswa/blackjack-fsm/player"
)

func TestNewPlayer(t *testing.T) {
	player := player.NewPlayer("Alice", player.WithInitialBankroll(500))

	if player.Name != "Alice" {
		t.Errorf("Expected player name to be 'Alice', got '%s'", player.Name)
	}

	if player.Bankroll != 500 {
		t.Errorf("Expected initial bankroll to be 500, got %d", player.Bankroll)
	}

	if len(player.Hands) != 0 {
		t.Errorf("Expected no hands, got %d", len(player.Hands))
	}
}
