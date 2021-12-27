package pass_the_bat

import (
	"fmt"
	"strings"
	"time"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
)

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
	passingBatTo := make(chan int, 1)
	var idx int = len(runningMen) - 1
	for idx >= 0 {
		go updateRunningManString(&runningMen[idx], idx, passingBatTo)
		idx = <-passingBatTo
	}
	*done = true
}

func updateRunningManString(runningMan *string, idx int, passingBatTo chan int) {
	stringSize := len([]rune(*runningMan)) - 1
	for i := stringSize; i > 0; i-- {
		time.Sleep(time.Second / 2)
		rmRuneSlice := []rune(*runningMan)
		rmRuneSlice[i], rmRuneSlice[i - 1] = rmRuneSlice[i - 1], rmRuneSlice[i]
		*runningMan = string(rmRuneSlice)
	}
	passingBatTo <- idx - 1
}