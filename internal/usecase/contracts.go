package usecase

import "github.com/andreyxaxa/tasks-api/internal/entity"

type (
	Tasks interface {
		CreateTask(entity.Task) entity.Task
		GetTask(string) (entity.Task, error)
		DeleteTask(string) error
	}
)
