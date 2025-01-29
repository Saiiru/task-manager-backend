package resolvers

import (
	"context"
	"fmt"
	"strconv"
	"task-manager-app/backend/internal/application"
	"task-manager-app/backend/internal/domain"
	"task-manager-app/backend/internal/interfaces/graphql/generated"
	"task-manager-app/backend/internal/interfaces/graphql/model"
	"time"
)

type Resolver struct {
	taskService *application.TaskService
	userService *application.UserService
}

func NewResolver(taskService *application.TaskService, userService *application.UserService) *Resolver {
	return &Resolver{
		taskService: taskService,
		userService: userService,
	}
}

// Root resolver implementations
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }
func (r *Resolver) Task() generated.TaskResolver         { return &taskResolver{r} }
func (r *Resolver) User() generated.UserResolver         { return &userResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
	taskResolver     struct{ *Resolver }
	userResolver     struct{ *Resolver }
)

// Task mutations
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*domain.Task, error) {
	userID, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	task := &domain.Task{
		Title:       input.Title,
		Description: input.Description, // Adicionando a descrição
		UserID:      userID,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.taskService.CreateTask(task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTask) (*domain.Task, error) {
	taskID, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID: %w", err)
	}

	task, err := r.taskService.GetTaskByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = *input.Description // Adicionando a descrição
	}
	if input.IsCompleted != nil {
		task.IsCompleted = *input.IsCompleted
	}
	task.UpdatedAt = time.Now()

	if err := r.taskService.UpdateTask(task); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}
	if err := r.taskService.DeleteTask(taskID); err != nil {
		return false, err
	}
	return true, nil
}

// Auth mutations
func (r *mutationResolver) Login(ctx context.Context, input model.UserLogin) (*domain.AuthResponse, error) {
	user, token, err := r.userService.Login(input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	return &domain.AuthResponse{User: user, Token: token}, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.UserRegister) (*domain.User, error) {
	user := &domain.User{
		Email:     input.Email,
		Name:      input.Name,
		LastName:  input.LastName,
		Avatar:    ptrStringValue(input.Avatar), // Adicionando o avatar
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.HashPassword(input.Password); err != nil {
		return nil, err
	}

	if err := r.userService.Register(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Query resolvers
func (r *queryResolver) Tasks(ctx context.Context, filter *model.TaskFilter) (*model.TaskConnection, error) {
	if filter == nil {
		filter = &model.TaskFilter{
			Page:  ptrInt(1),
			Limit: ptrInt(10),
		}
	}

	domainFilter := domain.TaskFilter{
		Search: ptrStringValue(filter.Search),
		Page:   ptrIntValue(filter.Page),
		Limit:  ptrIntValue(filter.Limit),
	}

	tasks, err := r.taskService.GetAllTasks(domainFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	edges := make([]*model.TaskEdge, len(tasks.Edges))
	for i, edge := range tasks.Edges {
		edges[i] = &model.TaskEdge{
			Node: &edge.Node,
		}
	}

	return &model.TaskConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     tasks.PageInfo.HasNextPage,
			HasPreviousPage: tasks.PageInfo.HasPreviousPage,
			TotalCount:      tasks.PageInfo.TotalCount,
		},
	}, nil
}

func (r *queryResolver) Task(ctx context.Context, id string) (*domain.Task, error) {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return r.taskService.GetTaskByID(taskID)
}

func (r *queryResolver) Me(ctx context.Context) (*domain.User, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	return r.userService.GetUserByID(id)
}

func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return r.userService.GetUserByID(userID)
}

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return r.userService.GetUserByEmail(email)
}

func (r *queryResolver) Users(ctx context.Context) ([]*domain.User, error) {
	return r.userService.GetAllUsers()
}

// Field resolvers
func (r *taskResolver) ID(ctx context.Context, obj *domain.Task) (string, error) {
	return strconv.Itoa(obj.ID), nil
}

func (r *taskResolver) UserID(ctx context.Context, obj *domain.Task) (string, error) {
	return strconv.Itoa(obj.UserID), nil
}

func (r *taskResolver) CreatedAt(ctx context.Context, obj *domain.Task) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *taskResolver) UpdatedAt(ctx context.Context, obj *domain.Task) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

func (r *taskResolver) Description(ctx context.Context, obj *domain.Task) (string, error) {
	return obj.Description, nil
}

func (r *userResolver) ID(ctx context.Context, obj *domain.User) (string, error) {
	return strconv.Itoa(obj.ID), nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *domain.User) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *domain.User) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) Avatar(ctx context.Context, obj *domain.User) (string, error) {
	if obj.Avatar == "" {
		return "", nil
	}
	return obj.Avatar, nil
}

// Helper functions
func ptrString(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}

func ptrStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ptrIntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
