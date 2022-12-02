package helpers

import (
	"fmt"
)

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func GetInfo(format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func GetWarning(format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
