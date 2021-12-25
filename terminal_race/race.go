package terminal_race

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
)

const racerRune rune = '🏎'

func StartRace(numberOfRacers int) {
	if numberOfRacers < 2 {
		panic("The number of racers must be greater than two.")
	}

	racers := make([]string, numberOfRacers)
	finished := make([]bool, numberOfRacers)
	racerLine := "______________🏎"

	for i := 0; i < numberOfRacers; i++ {
		racers[i] = racerLine
		printRacers(racers)
		go Update(&racers[i], finished, i)
	}

	for !go_concurrency_util.AnyTrue(finished) {
		printRacers(racers)
		time.Sleep(time.Millisecond)
		cmd := exec.Command("cmd", "/c", "cls")
		fmt.Fprintf(os.Stdout, "\x1b[?25l")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	fmt.Println(CheckRacer(finished))
}

func Update(racerLine *string, finished []bool, racerIdx int) {
	for !finished[racerIdx] {
		racerPos := strings.Index(*racerLine, string(racerRune))
		if racerPos - 1 >= 0 {
			runeSlice := []rune(*racerLine)
			runeSlice[racerPos], runeSlice[racerPos - 1] = runeSlice[racerPos - 1], runeSlice[racerPos]
			*racerLine = string(runeSlice)
		} else {
			finished[racerIdx] = true
		}
		time.Sleep(time.Duration(int(time.Millisecond) * rand.Intn(1000)))
	}
}

func CheckRacer(finished []bool) string {
	var result string
	for idx, racerFinished := range finished {
		if racerFinished {
			result = fmt.Sprintf("Racer %d won!", idx + 1)
		}
	}

	return result
}

func printRacers(racers []string) {
	for _, racer := range racers {
		fmt.Println(racer)
	}
}