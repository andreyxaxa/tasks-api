package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/andreyxaxa/tasks-api/internal/entity"
)

var (
	ErrNotFound = errors.New("task not found")
)

type TaskStorage struct {
	mu    sync.RWMutex
	tasks map[string]entity.Task
}

func New() *TaskStorage {
	return &TaskStorage{
		tasks: map[string]entity.Task{},
	}
}

func (s *TaskStorage) Create(t entity.Task) entity.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[t.ID] = t

	return t
}

func (s *TaskStorage) Get(id string) (entity.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	if !ok {
		return entity.Task{}, ErrNotFound
	}

	switch t.Status {
	case "In progress":
		t.InWork = time.Since(t.StartedAt)
	case "Done":
		t.InWork = t.Duration
	}

	return t, nil
}

func (s *TaskStorage) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}

	delete(s.tasks, id)

	return nil
}
