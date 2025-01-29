package unit

import (
	"task-manager-app/backend/internal/application"
	"task-manager-app/backend/internal/domain"
	"task-manager-app/backend/internal/infrastructure"
	"task-manager-app/backend/internal/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupUserService(t *testing.T) *application.UserService {
	db, err := tests.SetupTestDB()
	assert.NoError(t, err)

	repo := infrastructure.NewUserRepository(db)
	return application.NewUserService(repo)
}

func TestCreateUser(t *testing.T) {
	service := setupUserService(t)

	user := &domain.User{
		Name:         "John Doe",
		Email:        "john@example.com",
		PasswordHash: "hashedpassword",
	}
	err := service.Register(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUpdateUser(t *testing.T) {
	service := setupUserService(t)

	user := &domain.User{
		Name:         "John Doe",
		Email:        "john@example.com",
		LastName:     "Doe",
		Avatar:       "",
		PasswordHash: "hashedpassword",
	}
	err := service.Register(user)
	assert.NoError(t, err)

	user.Name = "Jane Doe"
	err = service.UpdateUser(user)
	assert.NoError(t, err)

	updatedUser, err := service.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", updatedUser.Name)
}

func TestDeleteUser(t *testing.T) {
	service := setupUserService(t)

	user := &domain.User{
		Name:         "John Doe",
		Email:        "john@example.com",
		LastName:     "Doe",
		Avatar:       "",
		PasswordHash: "hashedpassword",
	}

	err := service.Register(user)
	assert.NoError(t, err)

	err = service.DeleteUser(user.ID)
	assert.NoError(t, err)

	deletedUser, err := service.GetUserByID(user.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedUser)
}

func TestGetUser(t *testing.T) {
	service := setupUserService(t)

	user := &domain.User{
		Name:         "John Doe",
		Email:        "john@example.com",
		LastName:     "Doe",
		Avatar:       "",
		PasswordHash: "hashedpassword",
	}

	err := service.Register(user)
	assert.NoError(t, err)

	retrievedUser, err := service.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", retrievedUser.Name)
}
