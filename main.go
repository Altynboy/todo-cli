package main

import (
	"fmt"
	"todo-cli/file"
	"todo-cli/menu"
)

func main() {
	actions := []string{"list", "new", "open"}
	selIndex := 0

	selIndex, ok := menu.SelectOption(actions)
	if !ok {
		return
	}

	action := actions[selIndex]
	switch action {
	case "list":
		fmt.Println("List is printed")
	case "new":
		fmt.Println("New todo list")
		var filename string
		fmt.Println("Write todo name:")
		fmt.Scan(&filename)
		file.Create(filename)
	case "open":
		fmt.Println("Open todo list")
		var filename string
		fmt.Println("Write todo name:")
		fmt.Scan(&filename)
		file.Open(filename)
	}
}
