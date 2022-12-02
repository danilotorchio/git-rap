package helpers

import (
	"fmt"
	"os"
	"strings"
)

func CheckArgs(args ...string) {
	if len(os.Args) < len(args)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(args, " "))
		os.Exit(1)
	}
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func CheckIfEmpty(value, msg string) {
	if strings.TrimSpace(value) == "" {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", msg))
		os.Exit(1)
	}
}
