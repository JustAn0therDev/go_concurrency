package terminal_race

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
)

const racerRune rune = '🏎'

func StartRace(numberOfRacers int) {
	racers := make([]string, numberOfRacers)
	finished := make([]bool, numberOfRacers)
	racerLine := "______________🏎"

	for i := 0; i < numberOfRacers; i++ {
		racers[i] = racerLine
		go Update(&racers[i], finished, i)
	}

	for !go_concurrency_util.AnyTrue(finished) {
		printRacers(racers)
		time.Sleep(time.Microsecond)
		go_concurrency_util.ClearScreenAndHideCursor()
	}

	fmt.Printf("Racer %d won!\n", GetWinnerIndex(finished)+1)
}

func Update(racerLine *string, finished []bool, racerIdx int) {
	for !finished[racerIdx] {
		racerPos := strings.Index(*racerLine, string(racerRune))
		if racerPos-1 >= 0 {
			runeSlice := []rune(*racerLine)
			runeSlice[racerPos], runeSlice[racerPos-1] = runeSlice[racerPos-1], runeSlice[racerPos]
			*racerLine = string(runeSlice)
		} else {
			finished[racerIdx] = true
		}
		time.Sleep(time.Duration(int(time.Millisecond) * rand.Intn(700)))
	}
}

func printRacers(racers []string) {
	for _, racer := range racers {
		fmt.Println(racer)
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
