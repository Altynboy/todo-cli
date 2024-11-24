package main

import (
	"todo-cli/data"
	"todo-cli/handler"
	"todo-cli/menu"
)

func main() {
	options := []string{"list", "board", "new", "update", "delete"}
	selIndex := 0

	taskList := data.NewTodoData()
	taskList.Init()
	taskHandler := handler.Init(taskList)
	selIndex, ok := menu.SelectOption(options)
	if !ok {
		return
	}

	option := options[selIndex]
	switch option {
	case "list":
		taskHandler.DoPrint()
	case "board":
		taskHandler.DoBoard()
	case "new":
		taskHandler.DoNew()
	case "update":
		taskHandler.DoUpdate()
	case "delete":
		taskHandler.DoDelete()
	}
}
