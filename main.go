package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
	"github.com/JustAn0therDev/go_concurrency/pass_the_bat"
	"github.com/JustAn0therDev/go_concurrency/progress_bar"
	"github.com/JustAn0therDev/go_concurrency/terminal_race"
)

var demos = []string{"Terminal race", "Progress bar", "Pass the bat"}

var demoCalls = map[int]func() {
	0: startTerminalRace,
	1: startProgressBar,
	2: startPassTheBat,
}

// TODO: next would be nice to have some sort of "passing the baton" by using channels;
// much like runners passing a baton and only letting the next one run when the channel is
// triggered with a value.
// ____🏃____
// ________🏃 🏃________
// PASSED
// ________🏃 ________🏃
// ________🏃 _______🏃_
// ________🏃 ______🏃__
// ________🏃 _____🏃___
func main() {
	showAvailableDemos()

	demoIdx, err := getDemoIndex()

	go_concurrency_util.PanicIfErrNotNil(&err)

	demoCalls[demoIdx]()
}

func showAvailableDemos() {
	fmt.Println("Avaliable demos: ")

	for idx, demoName := range demos {
		fmt.Printf("%d - %v\n", idx, demoName)
	}

	fmt.Println()
}

func getDemoIndex() (int, error) {
	fmt.Print("Please type in the number of the demo you want to see: ")
	var input string

	_, err := fmt.Scanln(&input)

	if err != nil {
		return 0, err
	}

	demoIdx, err := strconv.ParseInt(input, 10, 32)

	if err != nil {
		return 0, err
	}

	if demoIdx < 0 || int(demoIdx) > len(demos) - 1 {
		return 0, errors.New("the demo index must be in the 'demos' range")
	}

	return int(demoIdx), nil
}

func startTerminalRace() {
	fmt.Print("Insert the number of cars you want to see in the race: ")
	var args string

	_, err := fmt.Scanln(&args)

	go_concurrency_util.PanicIfErrNotNil(&err)

	numberOfRacers, err := strconv.ParseInt(args, 10, 32)

	go_concurrency_util.PanicIfErrNotNil(&err)

	if numberOfRacers < 2 {
		panic("The number of racers must be at least 2 (two).")
	}

	terminal_race.StartRace(int(numberOfRacers))
}

func startProgressBar() {
	progress_bar.StartBar()
}

func startPassTheBat() {
	pass_the_bat.StartPassTheBat()
}