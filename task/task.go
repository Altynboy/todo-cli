package task

import (
	"fmt"
	"time"
)

type Task struct {
	Id        int
	Descr     string
	Status    Status
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func New(descr string, id int) (*Task, bool) {
	if descr == "" {
		return nil, false
	}
	now := time.Now()
	return &Task{
		Id:        id,
		Descr:     descr,
		Status:    StatusTodo,
		CreatedAt: now,
	}, true
}

func Update(task *Task, status Status) bool {
	if err := status.Validate(); err != nil {
		fmt.Println("Invalid status")
		return false
	}

	now := time.Now()
	task.UpdatedAt = &now
	task.Status = status
	return true
}
