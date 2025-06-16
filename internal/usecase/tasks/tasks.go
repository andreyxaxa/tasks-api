package tasks

import (
	"math/rand"
	"time"

	"github.com/andreyxaxa/tasks-api/internal/entity"
	"github.com/andreyxaxa/tasks-api/internal/repo"
	"github.com/google/uuid"
)

type UseCase struct {
	repo repo.TaskRepo
}

func New(r repo.TaskRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) CreateTask(t entity.Task) entity.Task {
	t.ID = uuid.NewString()
	t.Status = "Created"
	t.CreatedAt = time.Now()
	t.Duration = time.Duration(3+rand.Intn(3)) * time.Minute

	uc.repo.Create(t)

	go uc.runTask(t.ID)

	return t
}

func (uc *UseCase) GetTask(id string) (entity.Task, error) {
	t, err := uc.repo.Get(id)
	if err != nil {
		return entity.Task{}, err
	}

	t.InWork = time.Since(t.StartedAt)

	return t, nil
}

func (uc *UseCase) DeleteTask(id string) error {
	return uc.repo.Delete(id)
}

func (uc *UseCase) runTask(id string) {
	task, err := uc.repo.Get(id)
	if err != nil {
		return
	}

	task.Status = "In progress"
	task.StartedAt = time.Now()
	uc.repo.Create(task)

	// Симулируем работу задачи.
	time.Sleep(task.Duration)

	// По завершении работы, меняется статус и появляется "результат".
	task.Status = "Done"
	task.Result = "Completed successfully"
	uc.repo.Create(task)
}
