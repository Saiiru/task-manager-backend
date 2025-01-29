package graphql

import (
	"context"
	"task-manager-app/backend/internal/domain"
)

type TaskResolver interface {
	ID(ctx context.Context, obj *domain.Task) (string, error)
	UserID(ctx context.Context, obj *domain.Task) (string, error)
	CreatedAt(ctx context.Context, obj *domain.Task) (string, error)
	UpdatedAt(ctx context.Context, obj *domain.Task) (string, error)
}

type UserResolver interface {
	ID(ctx context.Context, obj *domain.User) (string, error)
	CreatedAt(ctx context.Context, obj *domain.User) (string, error)
	UpdatedAt(ctx context.Context, obj *domain.User) (string, error)
}
