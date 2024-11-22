package task

import "fmt"

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

func (s Status) Validate() error {
	switch s {
	case StatusTodo, StatusInProgress, StatusDone:
		return nil
	default:
		return fmt.Errorf("invalid status: %s", s)
	}
}

func StatusOptions() []string {
	return []string{string(StatusDone), string(StatusInProgress), string(StatusTodo)}
}
