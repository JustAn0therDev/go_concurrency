package pass_the_bat

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
)

func StartPassTheBat(competitors int) {
	var lines [][]string = make([][]string, competitors)
	var finished []bool = make([]bool, competitors)

	for i := 0; i < competitors; i++ {
		lines[i] = []string{"________🏃", "________🏃", "________🏃", "________🏃", "________🏃"}
		finished[i] = false
		go updateRunningMen(lines[i], &finished[i])
	}

	for !go_concurrency_util.AnyTrue(finished) {
		printLines(lines)
		time.Sleep(time.Microsecond)
		go_concurrency_util.ClearScreenAndHideCursor()
	}

	fmt.Printf("Line %d won!\n", GetWinnerIndex(finished) + 1)
}

func updateRunningMen(runningMen []string, done *bool) {
	passingBatTo := make(chan int, 1)
	runningManIdx := len(runningMen) - 1
	for runningManIdx >= 0 {
		go updateRunningManString(&runningMen[runningManIdx], runningManIdx, passingBatTo)
		runningManIdx = <-passingBatTo
	}
	*done = true
}

func updateRunningManString(runningMan *string, idx int, passingBatTo chan int) {
	stringSize := len([]rune(*runningMan)) - 1
	for i := stringSize; i > 0; i-- {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
		rmRuneSlice := []rune(*runningMan)
		rmRuneSlice[i], rmRuneSlice[i - 1] = rmRuneSlice[i - 1], rmRuneSlice[i]
		*runningMan = string(rmRuneSlice)
	}
	passingBatTo <- idx - 1
}

func printLines(lines [][]string) {
	for i := 0; i < len(lines); i++ {
		fmt.Println(strings.Join(lines[i], " "))
	}
}

func GetWinnerIndex(finished []bool) int {
	for idx, finish := range finished {
		if finish {
			return idx
		}
	}
	return -1
}