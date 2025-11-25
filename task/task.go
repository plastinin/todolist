package task

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	Title       string
	Description string
	TimeStart   time.Time
	Done        bool
	TimeEnd     time.Time
}

type TaskManager struct {
	TaskList map[string]*Task
}

func NewTaskManager() TaskManager {
	return TaskManager{
		TaskList: make(map[string]*Task, 0),
	}
}

func (tm TaskManager) Add(title string, description string) {
	tm.TaskList[title] = &Task{
		Title:       title,
		Description: description,
		TimeStart:   time.Now(),
		Done:        false,
	}
}

func (tm TaskManager) Delete(title string) {
	delete(tm.TaskList, title)
}

func (tm TaskManager) Done(title string) {
	if task, ok := tm.TaskList[title]; ok {
		task.Done = true
		task.TimeEnd = time.Now()
	}
}

func (tm TaskManager) PrintLn() {

	if len(tm.TaskList) == 0 {
		fmt.Println("Empty box. You are lucky")
		return
	}

	for _, v := range tm.TaskList {
		pp.Println(v)
	}
}
