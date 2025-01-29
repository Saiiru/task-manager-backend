package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task-manager-app/backend/internal/domain"
	"task-manager-app/backend/internal/interfaces"
	"task-manager-app/backend/internal/tests"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupUserDB() (*gorm.DB, error) {
	return tests.SetupTestDB()
}

func TestUserIntegration(t *testing.T) {
	db, err := setupUserDB()
	if err != nil {
		t.Fatalf("Failed to set up database: %v", err)
	}

	// Clear the users table before running the tests
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		t.Fatalf("Failed to clear users table: %v", err)
	}
	assert.NoError(t, err)
	router := interfaces.SetupRouter(db)

	t.Run("POST /register", func(t *testing.T) {
		user := domain.UserRegister{Name: "John Doe", LastName: "Marshal", Avatar: "", Email: "john@example.com", Password: "password"}
		body, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer valid_token")
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("GET /users", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users", nil)
		req.Header.Set("Authorization", "Bearer valid_token")
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("PUT /users/:id", func(t *testing.T) {
		user := domain.User{Name: "Jane Doe", Email: "jane@example.com", LastName: "Doe", Avatar: ""}
		body, _ := json.Marshal(user)

		req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer valid_token")
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("DELETE /users/:id", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/users/1", nil)
		req.Header.Set("Authorization", "Bearer valid_token")
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})
}
