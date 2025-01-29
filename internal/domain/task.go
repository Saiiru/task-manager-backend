package domain

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"` // Adicionando a descrição
	IsCompleted bool      `json:"isCompleted"`
	UserID      int       `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type NewTask struct {
	Title       string `json:"title"`
	Description string `json:"description"` // Adicionando a descrição
	UserID      string `json:"userId"`
}

type UpdateTask struct {
	ID          string  `json:"id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"` // Adicionando a descrição
	IsCompleted *bool   `json:"isCompleted,omitempty"`
}

type TaskFilter struct {
	Search string `json:"search"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	UserID string `json:"userId"`
}

type TaskEdge struct {
	Node Task `json:"node"`
}

type TaskConnection struct {
	Edges    []TaskEdge `json:"edges"`
	PageInfo struct {
		HasNextPage     bool `json:"hasNextPage"`
		HasPreviousPage bool `json:"hasPreviousPage"`
		TotalCount      int  `json:"totalCount"`
	} `json:"pageInfo"`
}

type TaskRepository interface {
	Create(task *Task) error
	FindByID(id int) (*Task, error)
	FindAll(filter TaskFilter) (*TaskConnection, error)
	Update(task *Task) error
	Delete(id int) error
	FindByUserID(userID int) ([]Task, error)
}
