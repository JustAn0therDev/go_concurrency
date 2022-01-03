package tests

import (
	"testing"
	"github.com/JustAn0therDev/go_concurrency/terminal_race"
)


func TestGetWinnerIndexTerminalRace(t *testing.T) {
	constWinnerIdx := 0
	finished := make([]bool, 10)
	finished[constWinnerIdx] = true
	winnerIdx := terminal_race.GetWinnerIndex(finished)
	if winnerIdx != 0 {
		t.Errorf("Expected %d, got %d", constWinnerIdx, winnerIdx)
	}
}