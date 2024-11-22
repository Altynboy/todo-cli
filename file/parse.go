package file

import (
	"encoding/json"
	"log"
	"todo-cli/task"
)

func Parse(file *[]byte) *[]task.Task {
	var tasks []task.Task
	if err := json.Unmarshal(*file, &tasks); err != nil {
		log.Fatal("Can't parse file", err)
	}

	return &tasks
}
