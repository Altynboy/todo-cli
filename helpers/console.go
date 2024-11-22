package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// clearConsole clears the terminal screen
func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Can't read user input %s", err)
	}
	return string(input[:len(input)-1])
}
