package terminal_race

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const racerRune rune = '🏎'

func StartRace() {
	racerOne := "______________🏎"
	racerTwo := "______________🏎"
	oneFinished := false
	twoFinished := false

	go Update(&racerOne, &oneFinished)
	go Update(&racerTwo, &twoFinished)

	for !oneFinished && !twoFinished {
		fmt.Println(racerOne)
		fmt.Println(racerTwo)
		time.Sleep(time.Millisecond)
		cmd := exec.Command("cmd", "/c", "cls")
		fmt.Fprintf(os.Stdout, "\x1b[?25l")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	fmt.Println(CheckRacer(oneFinished))
}

func Update(racerLine *string, finished *bool) {
	for !*finished {
		racerPos := strings.Index(*racerLine, string(racerRune))
		if racerPos - 1 >= 0 {
			runeSlice := []rune(*racerLine)
			runeSlice[racerPos], runeSlice[racerPos - 1] = runeSlice[racerPos - 1], runeSlice[racerPos]
			*racerLine = string(runeSlice)
		} else {
			*finished = true
		}
		time.Sleep(time.Duration(int(time.Millisecond) * rand.Intn(1000)))
	}
}

func CheckRacer(racerOneWon bool) string {
	if racerOneWon {
		return "1 won!"
	}

	return "2 won!"
}
