package tests

import (
	"testing"
	"github.com/JustAn0therDev/go_concurrency/pass_the_bat"
)

func TestGetWinnerIndexBat(t *testing.T) {
	constWinnerIdx := 0
	finished := make([]bool, 10)
	finished[constWinnerIdx] = true
	winnerIdx := pass_the_bat.GetWinnerIndex(finished)
	if winnerIdx != 0 {
		t.Errorf("Expected %d, got %d", constWinnerIdx, winnerIdx)
	}
}