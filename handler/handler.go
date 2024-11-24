package handler

import (
	"fmt"
	"log"
	"todo-cli/data"
	"todo-cli/helpers"
	"todo-cli/menu"
	"todo-cli/task"
)

type Handler struct {
	taskList *data.TodoData
}

func Init(taskList *data.TodoData) *Handler {
	return &Handler{
		taskList: taskList,
	}
}

func (h *Handler) DoPrint() {
	h.taskList.Print()
}

func (h *Handler) DoBoard() {
	h.taskList.Board()
}

func (h *Handler) DoNew() {
	fmt.Println("New todo")
	fmt.Println("Write todo descr:")
	descr := helpers.ReadLine()
	if ok := h.taskList.Write(descr); !ok {
		fmt.Println("cant write new task")
	}
	fmt.Println("New task successfully created")
	h.taskList.Read()
	h.taskList.Print()
}

func (h *Handler) DoUpdate() {
	fmt.Println("Chose todo to update status:")
	options := h.taskList.Options()
	selIndex, ok := menu.SelectOption(*options)
	if !ok {
		return
	}
	statusList := task.StatusOptions()
	selStatus, ok := menu.SelectOption(statusList)
	if !ok {
		return
	}
	ok = h.taskList.Update(selIndex, task.Status(statusList[selStatus]))
	if !ok {
		log.Fatalf("Can't update status for todo with id %d", selIndex)
		return
	}
	fmt.Printf("Status %s succesfully setted for todo %q\n",
		statusList[selStatus], h.taskList.Tasks()[selIndex].Descr)
	h.taskList.Print()
}

func (h *Handler) DoDelete() {
deleteTask:
	fmt.Println("Chose todo to delete:")
	options := h.taskList.Options()
	selIndex, ok := menu.SelectOption(*options)
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
	ok = h.taskList.DeleteTask(selIndex)
	if !ok {
		log.Fatalf("Can't delete task with id %d", selIndex)
		return
	}
	fmt.Printf("Task with id %d successfully deleted", selIndex)
	h.taskList.Print()
}
