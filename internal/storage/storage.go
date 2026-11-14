package storage

import (
	"errors"
	"todo/internal/task"
)

var ErrNotFound = errors.New("Задача не найдена")
var ErrDuplicateID = errors.New("Задача с таким ID уже существует")

type MemoryStorage struct {
	tasks []*task.Task
	byID  map[int]*task.Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make([]*task.Task, 0),
		byID:  make(map[int]*task.Task),
	}

}

func (s *MemoryStorage) Add(t *task.Task) error {
	if _, exists := s.byID[t.ID()]; exists {
		return ErrDuplicateID
	}
	s.byID[t.ID()] = t
	s.tasks = append(s.tasks, t)

	return nil
}

func (s *MemoryStorage) GetByID(id int) (*task.Task, error) {
	if _, exists := s.byID[id]; !exists {
		return nil, ErrNotFound
	}
	return s.byID[id], nil
}

func (s *MemoryStorage) GetAll() []*task.Task {
	return s.tasks
}

func (s *MemoryStorage) Delete(id int) error {
	if _, exists := s.byID[id]; !exists {
		return ErrNotFound
	}
	delete(s.byID, id)
	for i, t := range s.tasks {
		if t.ID() == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			break
		}

	}
	return nil
}

func (s *MemoryStorage) FilterByStatus(status task.Status) []*task.Task {
	var result []*task.Task
	for _, t := range s.tasks {
		if t.Status() == status {
			result = append(result, t)

		}

	}
	return result
}

func (s *MemoryStorage) Update(t *task.Task) error {
	if _, exists := s.byID[t.ID()]; !exists {
		return ErrNotFound
	}
	s.byID[t.ID()] = t
	return nil
}
