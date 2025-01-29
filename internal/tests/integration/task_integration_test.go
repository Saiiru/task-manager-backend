package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"task-manager-app/backend/internal/domain"
	"task-manager-app/backend/internal/interfaces"
	"task-manager-app/backend/internal/tests"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTaskDB() (*gorm.DB, error) {
	return tests.SetupTestDB()
}

func TestCreateTask(t *testing.T) {
	db, err := setupTaskDB()
	assert.NoError(t, err)
	router := interfaces.SetupRouter(db)

	task := domain.Task{Title: "Test Task", Description: "Test Description"}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestGetTasks(t *testing.T) {
	db, err := setupTaskDB()
	assert.NoError(t, err)
	router := interfaces.SetupRouter(db)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdateTask(t *testing.T) {
	db, err := setupTaskDB()
	assert.NoError(t, err)
	router := interfaces.SetupRouter(db)

	task := domain.Task{Title: "Test Task", Description: "Test Description"}
	db.Create(&task)

	updatedTask := domain.Task{Title: "Updated Task", Description: "Updated Description"}
	taskJSON, _ := json.Marshal(updatedTask)

	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(task.ID), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteTask(t *testing.T) {
	db, err := setupTaskDB()
	assert.NoError(t, err)
	router := interfaces.SetupRouter(db)

	task := domain.Task{Title: "Test Task", Description: "Test Description"}
	db.Create(&task)

	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(task.ID), nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
}
