package tests

import (
	"testing"

	"github.com/JustAn0therDev/go_concurrency/progress_bar"
)

func TestProgressBar(t *testing.T) {
	done := false
	progressBar := "░░░░░░░░░░"

	progress_bar.UpdateProgressBar(&progressBar, &done)

	if !done {
		t.Error("expected the UpdateProgressBar function to run normally and update a bool address.\n")
	}
}