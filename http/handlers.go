package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"todolist/task"

	"github.com/gorilla/mux"
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
	
	var taskDTO TaskDTO

	// JSON parse
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		http.Error(w, NewErrorDTO(err.Error()).ToString(), http.StatusBadRequest)
		return
	}

	// Validate
	if err := taskDTO.ValidateForCreate(); err != nil {
		http.Error(w, NewErrorDTO(err.Error()).ToString(), http.StatusBadRequest)
		return
	}

	// Add task
	todoTask, err := h.todoList.AddTask(taskDTO.Title, taskDTO.Description);
	if err != nil {		
		errCode := http.StatusInternalServerError
		if errors.Is(err, task.ErrTaskAlreadyExist) {
			errCode = http.StatusConflict 
		}
		http.Error(w, NewErrorDTO(err.Error()).ToString(), errCode)
		return
	}

	// JSON answer
	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}

	// HTTP answer
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: 	/tasks/{title}
method: 	GET
info: 		pattern

succeed:
  - status code: 200 Ok
  - response body: JSON represented found task

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	
	title := mux.Vars(r)["title"]
	taskGet, err := h.todoList.GetTask(title)
	
	if err != nil {		
		errCode := http.StatusInternalServerError
		if errors.Is(err, task.ErrTaskNotFound) {
			errCode = http.StatusNotFound 
		}
		http.Error(w, NewErrorDTO(err.Error()).ToString(), errCode)
		return
	}

	// JSON answer
	b, err := json.MarshalIndent(taskGet, "", "    ")
	if err != nil {
		panic(err)
	}

	// HTTP answer
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: 	/tasks
method: 	GET
info: 		-

succeed:
  - status code: 200 Ok
  - response body: JSON represented found tasks

failed:
  - status code: 400, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	// HTTP answer
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: 	/tasks?completed=false
method: 	GET
info: 		query params

succeed:
  - status code: 200 Ok
  - response body: JSON represented found tasks

failed:
  - status code: 400, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	
	tasks := h.todoList.ListUncompletedTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	// HTTP answer
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: 	/tasks/{title}
method: 	PATCH
info: 		pattern + JSON in request body 

succeed:
  - status code: 200 Ok
  - response body: JSON represented changed tasks

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	
	var completeTaskDTO CompleteTaskDTO

	// JSON parse
	if err := json.NewDecoder(r.Body).Decode(&completeTaskDTO); err != nil {
		http.Error(w, NewErrorDTO(err.Error()).ToString(), http.StatusBadRequest)
		return
	}

	// pattern 
	title := mux.Vars(r)["title"]
	taskComplete, err := h.todoList.CompleteTask(title)
	
	if err != nil {		
		errCode := http.StatusInternalServerError
		if errors.Is(err, task.ErrTaskNotFound) {
			errCode = http.StatusNotFound 
		}
		http.Error(w, NewErrorDTO(err.Error()).ToString(), errCode)
		return
	}

	// JSON answer
	b, err := json.MarshalIndent(taskComplete, "", "    ")
	if err != nil {
		panic(err)
	}

	// HTTP answer
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}

}

/*
pattern: 	/tasks/{title}
method: 	DELETE
info: 		pattern

succeed:
  - status code: 204 No content
  - response body: -

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	
	title := mux.Vars(r)["title"]
	err := h.todoList.Delete(title)
	
	if err != nil {		
		errCode := http.StatusInternalServerError
		if errors.Is(err, task.ErrTaskNotFound) {
			errCode = http.StatusNotFound 
		}
		http.Error(w, NewErrorDTO(err.Error()).ToString(), errCode)
		return
	}

	// HTTP answer
	w.WriteHeader(http.StatusNoContent)
}
