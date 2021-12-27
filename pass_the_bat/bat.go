package pass_the_bat

import (
	"fmt"
	"strings"
	"time"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
)

// const runningManRune rune = '🏃'

// TODO: For now, there will only be a single line of men running and passing the batons.
// TODO: make a thread pass the bat to each running man with channels,
// but make it a race.
func StartPassTheBat() {
	done := false
	runningMen := []string{"________🏃", "________🏃", "________🏃", "________🏃", "________🏃"}

	go UpdateRunningMen(runningMen, &done)

	for !done {
		fmt.Println(strings.Join(runningMen, " "))
		time.Sleep(time.Second)
		go_concurrency_util.ClearScreenAndHideCursor()
	}
}

func UpdateRunningMen(runningMen []string, done *bool) {
	passingBat := make(chan bool)
	for i := len(runningMen) - 1; i >= 0; i-- {
		go updateRunningManString(&runningMen[i], passingBat)
		<-passingBat
	}
	*done = true
}

func updateRunningManString(runningMan *string, passingBat chan bool) {
	stringSize := len([]rune(*runningMan)) - 1
	for i := stringSize; i > 0; i-- {
		time.Sleep(time.Second / 2)
		rmRuneSlice := []rune(*runningMan)
		rmRuneSlice[i], rmRuneSlice[i - 1] = rmRuneSlice[i - 1], rmRuneSlice[i]
		*runningMan = string(rmRuneSlice)
	}

	passingBat<-true
}