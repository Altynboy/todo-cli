package data

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"todo-cli/file"
	"todo-cli/helpers"
	"todo-cli/task"
)

type TodoData struct {
	tasks []task.Task
}

func NewTodoData() *TodoData {
	return &TodoData{}
}

var filename = "data"

func (t *TodoData) Init() {
	t.Read()
}

func (t *TodoData) Tasks() []task.Task {
	return t.tasks
}

func (t *TodoData) Read() {
	raw := file.Open(filename)
	t.tasks = *file.Parse(raw)
}

func (t *TodoData) Write(descr string) bool {
	data := t.tasks
	newId := data[len(data)-1].Id + 1
	if len(data) < 1 {
		newId = 0
	}
	new, ok := task.New(descr, newId)
	if !ok {
		log.Fatal("can't create new task")
	}
	data = append(data, *new)
	newData, err := json.MarshalIndent(data, "", "  ")
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

func (t *TodoData) Update(id int, status task.Status) bool {
	ok := task.Update(&(t.tasks)[id], status)
	if !ok {
		log.Fatal("cant update todo status")
		return false
	}

	newData, err := json.MarshalIndent(t.tasks, "", "  ")
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

func (t *TodoData) Print() {
	if len(t.tasks) < 1 {
		fmt.Println("Todo list is empty")
	}
	fmt.Printf("   %-20s \t%s\n", "TITLE", "STATUS")
	fmt.Println(strings.Repeat("-", 40))
	for i, task := range t.tasks {
		fmt.Printf("%d. %-20s \t%s\n", i+1, task.Descr, task.Status)
	}
}

func (t *TodoData) Options() *[]string {
	if len(t.tasks) < 1 {
		fmt.Println("Todo list is empty")
	}
	options := make([]string, len(t.tasks))
	for i, task := range t.tasks {
		options[i] = fmt.Sprintf("%d. %s \t%s", i+1, task.Descr, task.Status)
	}
	return &options
}

func (t *TodoData) Board() {
	var todoTasks []task.Task
	var inProgressTasks []task.Task
	var doneTasks []task.Task

	for _, el := range t.tasks {
		switch el.Status {
		case task.StatusTodo:
			todoTasks = append(todoTasks, el)
		case task.StatusInProgress:
			inProgressTasks = append(inProgressTasks, el)
		case task.StatusDone:
			doneTasks = append(doneTasks, el)
		}
	}

	max := helpers.Max(len(todoTasks), len(inProgressTasks))
	max = helpers.Max(max, len(doneTasks))

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
			fmt.Printf("%-*s ", 20, helpers.TrancuteOrWrap(col, 20))
		}
		fmt.Println()
	}
}
