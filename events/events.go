package events

import (
	"fmt"
	"time"
)

type Event struct {
	Command string
	Error   string
	Time    time.Time
}

type EventManager struct {
	List []Event
}

func NewEventManager() *EventManager {
	return &EventManager{
		List: make([]Event, 0),
	}
}

func (em *EventManager) Add(command string, Error string) {
	em.List = append(em.List, Event{
		Command: command,
		Error:   Error,
		Time:    time.Now(),
	})
}

func (em *EventManager) Println() {
	for v := range em.List {
		el := em.List[v]
		fmt.Printf("%d. %s %s %s\n", v+1, el.Time.Format("02.01.2006 в 15:04:05"), el.Command, el.Error)
	}
}
