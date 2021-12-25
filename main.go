package main

import (
	"os"
	"strconv"

	"github.com/JustAn0therDev/go_concurrency/terminal_race"
)

func main() {
	numberOfRacers, err := strconv.ParseInt(os.Args[1], 10, 32)

	if err != nil {
		panic(err)
	}

	terminal_race.StartRace(int(numberOfRacers))
}