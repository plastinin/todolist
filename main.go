package main

import (
	"fmt"
	"todolist/events"
	"todolist/http"
	"todolist/scanner"
	"todolist/task"
)

func main() {

	tm := task.NewTaskManager();
	
	server := http.NewHTTPServer(http.NewHTTPHandlers(tm))
	go server.Start()
	fmt.Println("HTTP-Server started")

	scanner := scanner.NewScanner(tm, events.NewEventManager())
	scanner.Start()
}
