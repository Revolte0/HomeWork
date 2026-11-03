package storage

import "todo/internal/task"

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
