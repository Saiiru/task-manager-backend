package infrastructure

import (
	"fmt"
	"task-manager-app/backend/internal/domain"
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *domain.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	if err := r.db.Create(task).Error; err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (r *TaskRepository) FindByID(id int) (*domain.Task, error) {
	var task domain.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, fmt.Errorf("failed to find task: %w", err)
	}
	return &task, nil
}

func (r *TaskRepository) FindAll(filter domain.TaskFilter) (*domain.TaskConnection, error) {
	var tasks []domain.Task
	var count int64
	query := r.db.Model(&domain.Task{})

	if filter.Search != "" {
		query = query.Where("title LIKE ?", "%"+filter.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, fmt.Errorf("failed to count tasks: %w", err)
	}

	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Offset(offset).Limit(filter.Limit)
	}

	if err := query.Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("failed to find tasks: %w", err)
	}

	var edges []domain.TaskEdge
	for _, task := range tasks {
		edges = append(edges, domain.TaskEdge{Node: task})
	}
	connection := &domain.TaskConnection{
		Edges: edges,
	}
	connection.PageInfo.TotalCount = int(count)
	connection.PageInfo.HasNextPage = len(tasks) == filter.Limit
	connection.PageInfo.HasPreviousPage = filter.Page > 1

	return connection, nil
}

func (r *TaskRepository) Update(task *domain.Task) error {
	task.UpdatedAt = time.Now()
	if err := r.db.Save(task).Error; err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (r *TaskRepository) Delete(id int) error {
	if err := r.db.Delete(&domain.Task{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (r *TaskRepository) FindByUserID(userID int) ([]domain.Task, error) {
	var tasks []domain.Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("failed to find tasks by user ID: %w", err)
	}
	return tasks, nil
}
