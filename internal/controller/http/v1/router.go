package v1

import (
	"net/http"

	"github.com/andreyxaxa/tasks-api/internal/usecase"
	"github.com/gorilla/mux"
)

func NewTasksRoutes(apiV1Group *mux.Router, t usecase.Tasks) {
	r := &V1{
		t: t,
	}

	tasksGroup := apiV1Group.PathPrefix("/tasks").Subrouter()

	tasksGroup.HandleFunc("", r.create).Methods(http.MethodPost)
	tasksGroup.HandleFunc("/{id}", r.get).Methods(http.MethodGet)
	tasksGroup.HandleFunc("/{id}", r.delete).Methods(http.MethodPost)
}
