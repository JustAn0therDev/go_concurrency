package main

import (
	"os"
	"strconv"

	"github.com/JustAn0therDev/go_concurrency/terminal_race"
)

func main() {
	if len(os.Args) != 2 {
		panic("The number of racers must be provided and be the only argument to the program.")
	}

	numberOfRacers, err := strconv.ParseInt(os.Args[1], 10, 32)

	if err != nil {
		panic(err)
	}

	terminal_race.StartRace(int(numberOfRacers))
}