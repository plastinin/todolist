package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todolist/events"
	"todolist/task"
)

type Scanner struct {
	TaskManager *task.TaskManager
	EventManager *events.EventManager
}

func NewScanner(TaskManager *task.TaskManager, EventManager *events.EventManager) *Scanner {
	return &Scanner{
		TaskManager: TaskManager,
		EventManager: EventManager,	
	}
}

func (sc *Scanner) Start() {

	printAbout()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printEnterComand()
		
		if ok := scanner.Scan(); !ok {
			printError(errorInput)
			continue
		}

		command := scanner.Text()
		commands := strings.SplitN(command, " ", 3)

		if len(commands) == 0 {
			printError(errorInput)
			continue
		}

		switch commands[0] {
		case "/help":
			printHelp()
		case "/add":
			sc.TaskManager.Add(commands[1], commands[2])
			fmt.Println("Done.")
		case "/list":
			sc.TaskManager.PrintLn()
		case "/del":
			sc.TaskManager.Delete(commands[1])
			fmt.Println("Done.")
		case "/done":
			sc.TaskManager.DoneTask(commands[1])
			fmt.Println("Done.")
		case "/events":
			sc.EventManager.Println()
		case "/exit":
			return
		default:
			printError(errorNoComand)
		}
		sc.EventManager.Add(command, "")
	}
}

