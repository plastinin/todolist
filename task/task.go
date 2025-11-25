package task

import (
	"time"
)

type Task struct {
	Title       string
	Description string
	TimeStart   time.Time
	IsDone        bool
	TimeEnd     *time.Time
}

func NewTask(title string, desc string) *Task {
	return &Task{
		Title: title,
		Description: desc,
		TimeStart: time.Now(),
		IsDone: false,
		TimeEnd: nil,
	}
}

func (t *Task) Done() {
	t.IsDone = true
	now := time.Now()
	t.TimeEnd = &now
}