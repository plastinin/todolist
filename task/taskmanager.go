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

func (tm *TaskManager) AddTask(title string, description string) error {

	if _, ok := tm.TaskList[title]; ok {
		return ErrTaskAlreadyExist
	}
	tm.TaskList[title] = NewTask(title, description)
	return nil
}

func (tm *TaskManager) Delete(title string) error {

	if _, ok := tm.TaskList[title]; !ok {
		return ErrTaskNotFound
	}

	delete(tm.TaskList, title)
	return nil
}

func (tm *TaskManager) CompleteTask(title string) error {
	if task, ok := tm.TaskList[title]; ok {
		task.Complete()
		return nil
	} else {
		return ErrTaskNotFound
	}
}

func (tm *TaskManager) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(tm.TaskList))
	for k, v := range tm.TaskList {
		tmp[k] = *v
	}
	return tmp
}

func (tm *TaskManager) ListNotCompletedTasks() map[string]Task {
	tmp := make(map[string]Task, len(tm.TaskList))
	for k, v := range tm.TaskList {
		if v.Completed {
			tmp[k] = *v
		}
	}
	return tmp
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
