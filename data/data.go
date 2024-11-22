package data

import (
	"encoding/json"
	"fmt"
	"log"
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
