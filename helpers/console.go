package helpers

import "fmt"

// clearConsole clears the terminal screen
func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}
