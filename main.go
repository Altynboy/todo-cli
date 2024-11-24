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
	actions := []string{"list", "board", "new", "update", "delete"}
	selIndex := 0

	taskList := data.NewTodoData()
	taskList.Init()
	selIndex, ok := menu.SelectOption(actions)
	if !ok {
		return
	}

	action := actions[selIndex]
	switch action {
	case "list":
		taskList.Print()
	case "new":
		fmt.Println("New todo")
		fmt.Println("Write todo descr:")
		descr := helpers.ReadLine()
		if ok := taskList.Write(descr); !ok {
			fmt.Println("cant write new task")
		}
		fmt.Println("New task successfully created")
		taskList.Read()
		taskList.Print()
	case "update":
		fmt.Println("Chose todo to update status:")
		options := taskList.Options()
		selIndex, ok = menu.SelectOption(*options)
		if !ok {
			return
		}
		statusList := task.StatusOptions()
		selStatus, ok := menu.SelectOption(statusList)
		if !ok {
			return
		}
		ok = taskList.Update(selIndex, task.Status(statusList[selStatus]))
		if !ok {
			log.Fatalf("Can't update status for todo with id %d", selIndex)
			return
		}
		fmt.Printf("Status %s succesfully setted for todo %q\n", statusList[selStatus], taskList.Tasks()[selIndex].Descr)
		taskList.Print()
	case "board":
		taskList.Board()
	case "delete":
	deleteTask:
		fmt.Println("Chose todo to delete:")
		options := taskList.Options()
		selIndex, ok = menu.SelectOption(*options)
		if !ok {
			return
		}
		confirm := []string{"yes", "no"}
		selConfirm, ok := menu.SelectOption(confirm)
		if !ok {
			return
		}
		if selConfirm == 1 {
			goto deleteTask
		}
		ok = taskList.DeleteTask(selIndex)
		if !ok {
			log.Fatalf("Can't delete task with id %d", selIndex)
			return
		}
		fmt.Printf("Task with id %d successfully deleted", selIndex)
		taskList.Print()
	}
}
