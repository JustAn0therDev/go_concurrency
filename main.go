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
	1: progress_bar.StartBar,
	2: startPassTheBat,
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

func startPassTheBat() {
	fmt.Print("Insert the number of lines you want to see in the competition: ")
	var args string

	_, err := fmt.Scanln(&args)

	go_concurrency_util.PanicIfErrNotNil(&err)

	numberOfLines, err := strconv.ParseInt(args, 10, 32)

	go_concurrency_util.PanicIfErrNotNil(&err)

	if numberOfLines < 2 {
		panic("The number of lines must be at least 2 (two).")
	}

	pass_the_bat.StartPassTheBat(int(numberOfLines))
}