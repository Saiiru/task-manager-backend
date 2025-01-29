package domain

import (
	"errors"
	"regexp"
	"task-manager-app/backend/pkg/utils"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Name         string    `json:"name"`
	LastName     string    `json:"lastName"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Avatar   string `json:"avatar"`
}

type AuthResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type UserUpdate struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Avatar   string `json:"avatar"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.PasswordHash = hashedPassword
	return nil
}

func (u *User) CheckPassword(password string) error {
	if !utils.CheckPasswordHash(password, u.PasswordHash) {
		return errors.New("invalid password")
	}
	return nil
}

func (u *UserRegister) Validate() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindAll() ([]User, error)
	FindByID(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
}
