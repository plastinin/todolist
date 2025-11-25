package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todolist/events"
	"todolist/task"

	"github.com/k0kubun/pp"
)

func main() {
	tm := task.NewTaskManager()
	em := events.NewEventManager()

	fmt.Println("Task manager v0.1")
	fmt.Println("=================\n\n")

	for {
		fmt.Println("Enter a command to continue. Hint: /help")
		scanner := bufio.NewScanner(os.Stdin)

		if ok := scanner.Scan(); !ok {
			pp.Println("Error input. Please try again...")
			continue
		}

		command := scanner.Text()
		commands := strings.SplitN(command, " ", 3)

		switch commands[0] {
		case "/help":
			printHelp()
		case "/add":
			tm.Add(commands[1], commands[2])
			fmt.Println("Done.")
		case "/list":
			tm.PrintLn()
		case "/del":
			tm.Delete(commands[1])
			fmt.Println("Done.")
		case "/done":
			tm.Done(commands[1])
			fmt.Println("Done.")
		case "/events":
			em.Println()
		case "/exit":
			return
		}
		em.Add(command, "")
	}
}

func printHelp() {
	fmt.Println(`Список команд, которые должны быть доступны в приложении:
		/help — эта команда позволяет узнать доступные команды и их формат
		/add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач
		/list — эта команда позволяет получить полный список всех задач
		/del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку
		/done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную
		/events — эта команда позволяет получить список всех событий
		/exit — эта команда позволяет завершить выполнение программы`)
}
