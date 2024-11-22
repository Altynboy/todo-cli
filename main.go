package main

import (
	"fmt"
	"log"
	"todo-cli/data"
	"todo-cli/helpers"
	"todo-cli/menu"
	"todo-cli/task"
)

func main() {
	actions := []string{"list", "new", "update"}
	selIndex := 0

	taskList := data.Read()

	selIndex, ok := menu.SelectOption(actions)
	if !ok {
		return
	}

	action := actions[selIndex]
	switch action {
	case "list":
		data.Print(taskList)
	case "new":
		fmt.Println("New todo")
		fmt.Println("Write todo descr:")
		descr := helpers.ReadLine()
		if ok := data.Write(taskList, descr); !ok {
			fmt.Println("cant write new task")
		}
		fmt.Println("New task successfully created")
		newTaskList := data.Read()
		data.Print(newTaskList)
	case "update":
		fmt.Println("Chose todo to update status:")
		options := data.Options(taskList)
		selIndex, ok = menu.SelectOption(*options)
		if !ok {
			return
		}
		statusList := task.StatusOptions()
		selStatus, ok := menu.SelectOption(statusList)
		if !ok {
			return
		}
		ok = data.Update(taskList, selIndex, task.Status(statusList[selStatus]))
		if !ok {
			log.Fatalf("Can't update status for todo with id %d", selIndex)
			return
		}
		fmt.Printf("Status %s succesfully setted for todo %q\n", statusList[selStatus], (*taskList)[selIndex].Descr)
		data.Print(taskList)
	}
}
