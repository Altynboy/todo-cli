package data

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"todo-cli/file"
	"todo-cli/task"
)

var filename = "data"

func Read() *[]task.Task {
	raw := file.Open(filename)
	return file.Parse(raw)
}

func Write(data *[]task.Task, descr string) bool {
	newId := (*data)[len(*data)-1].Id + 1
	if len(*data) < 1 {
		newId = 0
	}
	new, ok := task.New(descr, newId)
	if !ok {
		log.Fatal("can't create new task")
	}
	*data = append(*data, *new)
	newData, err := json.MarshalIndent(*data, "", "  ")
	if err != nil {
		log.Fatalf("error while marshalling new data %s", err)
		return false
	}

	if err := file.Write(filename, &newData); err != nil {
		log.Fatalf("error while writing to file %s", err)
		return false
	}

	return true
}

func Update(data *[]task.Task, id int, status task.Status) bool {
	ok := task.Update(&(*data)[id], status)
	if !ok {
		log.Fatal("cant update todo status")
		return false
	}

	newData, err := json.MarshalIndent(*data, "", "  ")
	if err != nil {
		log.Fatalf("error while marshalling new data %s", err)
		return false
	}

	if err := file.Write(filename, &newData); err != nil {
		log.Fatalf("error while writing to file %s", err)
		return false
	}

	return true
}

func Print(tasks *[]task.Task) {
	if len(*tasks) < 1 {
		fmt.Println("Todo list is empty")
	}
	for i, task := range *tasks {
		fmt.Printf("%d. %s \t%s\n", i+1, task.Descr, task.Status)
	}
}

func Options(tasks *[]task.Task) *[]string {
	if len(*tasks) < 1 {
		fmt.Println("Todo list is empty")
	}
	options := make([]string, len(*tasks))
	for i, task := range *tasks {
		options[i] = fmt.Sprintf("%d. %s \t%s", i+1, task.Descr, task.Status)
	}
	return &options
}

func Board(tasks *[]task.Task) {
	var todoTasks []task.Task
	var inProgressTasks []task.Task
	var doneTasks []task.Task

	for _, el := range *tasks {
		switch el.Status {
		case task.StatusTodo:
			todoTasks = append(todoTasks, el)
		case task.StatusInProgress:
			inProgressTasks = append(inProgressTasks, el)
		case task.StatusDone:
			doneTasks = append(doneTasks, el)
		}
	}

	max := Max(len(todoTasks), len(inProgressTasks))
	max = Max(max, len(doneTasks))

	headers := task.StatusOptions()
	fmt.Printf("   %-20s %-20s %-20s\n", headers[0], headers[1], headers[2])
	fmt.Println(strings.Repeat("-", 60))
	for i := 0; i < max; i++ {
		row := make([]string, len(headers))
		if i < len(todoTasks) {
			row[0] = todoTasks[i].Descr
		}
		if i < len(inProgressTasks) {
			row[1] = inProgressTasks[i].Descr
		}
		if i < len(doneTasks) {
			row[2] = doneTasks[i].Descr
		}
		fmt.Printf("%d. ", i+1)

		for _, col := range row {
			fmt.Printf("%-*s ", 20, TrancuteOrWrap(col, 20))
		}
		fmt.Println()
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func TrancuteOrWrap(text string, width int) string {
	if len(text) <= width {
		return text
	}

	var wrapped string
	for len(text) > width {
		wrapped += text[:width] + "\n"
		text = text[width:]
	}
	wrapped += text

	return wrapped
}
