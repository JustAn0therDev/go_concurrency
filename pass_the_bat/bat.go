package pass_the_bat

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
)

func StartPassTheBat(competitors int) {
	var lines [][]string
	var finished []bool

	for i := 0; i < competitors; i++ {
		lines = append(lines, []string{"________🏃", "________🏃", "________🏃", "________🏃", "________🏃"})
		finished = append(finished, false)
		go updateRunningMen(lines[i], &finished[i])
	}

	for !go_concurrency_util.AnyTrue(finished) {
		printLines(lines)
		time.Sleep(time.Millisecond)
		go_concurrency_util.ClearScreenAndHideCursor()
	}

	printWinner(finished)
}

func updateRunningMen(runningMen []string, done *bool) {
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
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
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

func printWinner(finished []bool) {
	for idx, finish := range finished {
		if finish {
			fmt.Printf("Line %d won!\n", idx + 1)
			break
		}
	}
}