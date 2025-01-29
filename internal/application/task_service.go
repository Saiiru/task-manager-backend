package application

import (
	"task-manager-app/backend/internal/domain"
)

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *domain.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) GetTaskByID(id int) (*domain.Task, error) {
	return s.repo.FindByID(id)
}

func (s *TaskService) GetAllTasks(filter domain.TaskFilter) (*domain.TaskConnection, error) {
	return s.repo.FindAll(filter)
}

func (s *TaskService) UpdateTask(task *domain.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}

func (s *TaskService) GetTasksByUserID(userID int) ([]domain.Task, error) {
	return s.repo.FindByUserID(userID)
}
