package go_concurrency_util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreenAndHideCursor() {
	runningOs := runtime.GOOS
	var cmd *exec.Cmd
	switch runningOs {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	fmt.Fprintf(os.Stdout, "\x1b[?25l")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
