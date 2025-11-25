package task

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type TaskManager struct {
	TaskList map[string]*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		TaskList: make(map[string]*Task, 0),
	}
}

func (tm *TaskManager) Add(title string, description string) {
	tm.TaskList[title] = NewTask(title, description)
}

func (tm *TaskManager) Delete(title string) {
	delete(tm.TaskList, title)
}

func (tm *TaskManager) DoneTask(title string) {
	if task, ok := tm.TaskList[title]; ok {
		task.Done()
	}
}

func (tm *TaskManager) PrintLn() {

	if len(tm.TaskList) == 0 {
		fmt.Println("Empty box. You are lucky")
		return
	}

	for _, v := range tm.TaskList {
		pp.Println(v)
	}
}
