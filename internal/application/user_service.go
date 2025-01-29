package application

import (
	"errors"
	"strconv"
	"task-manager-app/backend/internal/domain"
	"task-manager-app/backend/pkg/utils"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *domain.User) error {
	// Verificar se o email já existe
	existingUser, err := s.repo.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already in use")
	}

	if err := utils.ValidatePassword(user.PasswordHash); err != nil {
		return err
	}
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (*domain.User, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, "", errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(strconv.Itoa(user.ID))
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	// Convert []domain.User to []*domain.User
	userPtrs := make([]*domain.User, len(users))
	for i, user := range users {
		userPtrs[i] = &user
	}

	return userPtrs, nil
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
