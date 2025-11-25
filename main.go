package main

import (
	"todolist/events"
	"todolist/scanner"
	"todolist/task"
)

func main() {
	scanner := scanner.NewScanner(
		task.NewTaskManager(),
		events.NewEventManager(),
	)
	scanner.Start()
}
