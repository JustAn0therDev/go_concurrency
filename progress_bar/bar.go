package progress_bar

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
)

var loadedprogressBarRune rune = '▓'

func StartBar() {
	done := false
	progressBar := "░░░░░░░░░░"

	go UpdateProgressBar(&progressBar, &done)

	for !done {
		fmt.Println(progressBar)
		go_concurrency_util.ClearScreenAndHideCursor()
	}
}

func UpdateProgressBar(progressBar *string, done *bool) {
	progressBarSize := len([]rune(*progressBar))
	for i := 0; i < progressBarSize; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		pbRuneSlice := []rune(*progressBar)
		pbRuneSlice[i] = loadedprogressBarRune
		*progressBar = string(pbRuneSlice)
	}
	*done = true
}