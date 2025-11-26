package http

import (
	"net/http"
	"todolist/task"
)

type HTTPHandlers struct {
	todoList *task.TaskManager
}

func NewHTTPHandlers(todolist *task.TaskManager) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todolist,
	}
}

/*
pattern: 	/tasks
method: 	POST
info: 		JSON in HTTP body

succeed:
  - status code: 201 Created
  - response body: JSON represent task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

}
