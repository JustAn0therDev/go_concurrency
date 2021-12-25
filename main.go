package main

import (
	"fmt"

	"github.com/JustAn0therDev/go_concurrency/concurrency_tests"
)

func main() {
	val := 1
	go concurrency_tests.IncrementAsync(&val)
	concurrency_tests.IncrementAsync(&val)
	fmt.Printf("Final value: %d\n", val)
}