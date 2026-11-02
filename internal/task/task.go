package task

import (
	"errors"
	"fmt"
	"time"
)

var nextID int = 1

type Status string

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	id          int       `json:"id"`
	title       string    `json:"title"`
	description string    `json:"description"`
	status      Status    `json:"status"`
	createdAt   time.Time `json:"createdAt"`
	updatedAt   time.Time `json:"updatedAt"`
}

func (t *Task) SetStatus(newStatus Status) {
	t.status = newStatus
	t.updatedAt = time.Now()
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %d | Название: %s | Описание: %s | Статус: %s", t.id, t.title, t.description, t.status)

}

func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("Ошибка: название не должно быть пустым")
	}
	task := &Task{
		id:          nextID,
		title:       title,
		description: description,
		status:      StatusNew,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
	nextID++

	return task, nil
}

func (t *Task) ID() int {
	return t.id
}
func (t *Task) Title() string {
	return t.title
}
func (t *Task) Description() string {
	return t.description
}
func (t *Task) Status() Status {
	return t.status
}
func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}
func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}
