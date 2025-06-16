package repo

import "github.com/andreyxaxa/tasks-api/internal/entity"

type (
	TaskRepo interface {
		Create(entity.Task) entity.Task
		Delete(string) error
		Get(string) (entity.Task, error)
	}
)
