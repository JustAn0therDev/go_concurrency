package go_concurrency_util

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreenAndHideCursor() {
	cmd := exec.Command("cmd", "/c", "cls")
	fmt.Fprintf(os.Stdout, "\x1b[?25l")
	cmd.Stdout = os.Stdout
	cmd.Run()
}