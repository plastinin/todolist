package task

import (
	"fmt"
	"sync"

	"github.com/k0kubun/pp"
)

type TaskManager struct {
	TaskList map[string]*Task
	mtx 	sync.RWMutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		TaskList: make(map[string]*Task, 0),
	}
}

func (tm *TaskManager) AddTask(title string, description string) (Task, error) {

	tm.mtx.Lock()
	defer tm.mtx.Unlock()

	if _, ok := tm.TaskList[title]; ok {
		return Task{}, ErrTaskAlreadyExist
	}

	task := NewTask(title, description)
	tm.TaskList[title] = task
	return *task, nil
}

func (tm *TaskManager) GetTask(title string) (Task, error) {

	tm.mtx.RLock()
	defer tm.mtx.RUnlock()

	task, ok := tm.TaskList[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return *task, nil
}


func (tm *TaskManager) Delete(title string) error {

	tm.mtx.Lock()
	defer tm.mtx.Unlock()

	if _, ok := tm.TaskList[title]; !ok {
		return ErrTaskNotFound
	}

	delete(tm.TaskList, title)
	return nil
}

func (tm *TaskManager) CompleteTask(title string) (Task, error) {

	tm.mtx.Lock()
	defer tm.mtx.Unlock()

	if task, ok := tm.TaskList[title]; ok {
		task.Complete()
		return *task, nil
	} else {
		return Task{}, ErrTaskNotFound
	}
}

func (tm *TaskManager) ListTasks() map[string]Task {

	tm.mtx.RLock()
	defer tm.mtx.RUnlock()

	tmp := make(map[string]Task, len(tm.TaskList))
	for k, v := range tm.TaskList {
		tmp[k] = *v
	}
	return tmp
}

func (tm *TaskManager) ListUncompletedTasks() map[string]Task {

	tm.mtx.RLock()
	defer tm.mtx.RUnlock()

	tmp := make(map[string]Task)
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
