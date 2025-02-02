// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"task-manager-app/backend/internal/domain"
)

type Mutation struct {
}

type NewTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}

type PageInfo struct {
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	TotalCount      int  `json:"totalCount"`
}

type Query struct {
}

type TaskConnection struct {
	Edges    []*TaskEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type TaskEdge struct {
	Node *domain.Task `json:"node"`
}

type TaskFilter struct {
	Search *string `json:"search,omitempty"`
	Page   *int    `json:"page,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type UpdateTask struct {
	ID          string  `json:"id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	IsCompleted *bool   `json:"isCompleted,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	LastName string  `json:"lastName"`
	Avatar   *string `json:"avatar,omitempty"`
}
