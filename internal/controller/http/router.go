package http

import (
	v1 "github.com/andreyxaxa/tasks-api/internal/controller/http/v1"
	"github.com/andreyxaxa/tasks-api/internal/usecase"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, t usecase.Tasks) {
	apiV1Group := r.PathPrefix("/v1").Subrouter()

	{
		v1.NewTasksRoutes(apiV1Group, t)
	}
}
