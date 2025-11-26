package task

import (
	"time"
)

type Task struct {
	Title       string
	Description string
	StartAt     time.Time
	Completed   bool
	CompletedAt *time.Time
}

func NewTask(title string, desc string) *Task {
	return &Task{
		Title:       title,
		Description: desc,
		StartAt:     time.Now(),
		Completed:   false,
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	t.Completed = true
	now := time.Now()
	t.CompletedAt = &now
}
