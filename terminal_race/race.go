package terminal_race

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var racerRuneOne rune = '●'
var racerRuneTwo rune = '☭'

func StartRace() {
	racerOne := "●--------------"
	racerTwo := "☭--------------"
	oneFinished := false
	twoFinished := false

	go Update(&racerOne, &racerRuneOne, &oneFinished)
	go Update(&racerTwo, &racerRuneTwo, &twoFinished)

	for !oneFinished && !twoFinished {
		fmt.Println(racerOne)
		fmt.Println(racerTwo)
		time.Sleep(time.Second / 2)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	fmt.Println(CheckRacer(oneFinished))
}

func Update(racerLine *string, racerRune *rune, finished *bool) {
	stringSize := len([]rune(*racerLine))
	for rune((*racerLine)[stringSize]) != *racerRune {
		racerPos := strings.Index(*racerLine, string(*racerRune))
		if racerPos + 1 < stringSize - 1 {
			runeSlice := []rune(*racerLine)
			runeSlice[racerPos] = '-'
			runeSlice[racerPos + 1] = rune(*racerRune)
			*racerLine = string(runeSlice)
		} else {
			*finished = true
			return
		}
		time.Sleep(time.Duration(int(time.Millisecond) * rand.Intn(1000)))
	}
}

func CheckRacer(racerOneWon bool) string {
	if racerOneWon {
		return "Racer one won!"
	}

	return "Racer two won!"
}
