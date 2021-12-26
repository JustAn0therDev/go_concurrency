package main

import (
	"fmt"
	"strconv"

	"github.com/JustAn0therDev/go_concurrency/go_concurrency_util"
	"github.com/JustAn0therDev/go_concurrency/terminal_race"
)

var demos = []string{"Terminal race"}

var demoCalls = map[int]func() {
	0: StartTerminalRace,
}

func main() {
	showAvailableDemos()

	demoIdx, err := getDemoIndex()

	go_concurrency_util.PanicIfErrNotNil(&err)

	demoCalls[demoIdx]()
}

func showAvailableDemos() {
	fmt.Println("Avaliable demos: ")

	for idx, demoName := range demos {
		fmt.Printf("%d - %v", idx, demoName)
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

	if demoIdx < 0 || int(demoIdx) > len(demos) {
		panic("The demo index must be in the 'demos' range.")
	}

	return int(demoIdx), nil
}

func StartTerminalRace() {
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