package unit

import (
	"task-manager-app/backend/internal/application"
	"task-manager-app/backend/internal/domain"
	"task-manager-app/backend/internal/infrastructure"
	"task-manager-app/backend/internal/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTaskService(t *testing.T) *application.TaskService {
	db, err := tests.SetupTestDB()
	assert.NoError(t, err)

	repo := infrastructure.NewTaskRepository(db)
	return application.NewTaskService(repo)
}

func TestCreateTask(t *testing.T) {
	service := setupTaskService(t)

	task := &domain.Task{
		Title:  "Test Task",
		UserID: 1,
	}
	err := service.CreateTask(task)
	assert.NoError(t, err)
	assert.NotZero(t, task.ID)
}

func TestUpdateTask(t *testing.T) {
	service := setupTaskService(t)

	task := &domain.Task{
		Title:  "Test Task",
		UserID: 1,
	}
	err := service.CreateTask(task)
	assert.NoError(t, err)

	task.Title = "Updated Task"
	err = service.UpdateTask(task)
	assert.NoError(t, err)

	updatedTask, err := service.GetTaskByID(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", updatedTask.Title)
}

func TestDeleteTask(t *testing.T) {
	service := setupTaskService(t)

	task := &domain.Task{
		Title:  "Test Task",
		UserID: 1,
	}
	err := service.CreateTask(task)
	assert.NoError(t, err)

	err = service.DeleteTask(task.ID)
	assert.NoError(t, err)

	deletedTask, err := service.GetTaskByID(task.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedTask)
}

func TestGetTask(t *testing.T) {
	service := setupTaskService(t)

	task := &domain.Task{
		Title:  "Test Task",
		UserID: 1,
	}
	err := service.CreateTask(task)
	assert.NoError(t, err)

	retrievedTask, err := service.GetTaskByID(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", retrievedTask.Title)
}
