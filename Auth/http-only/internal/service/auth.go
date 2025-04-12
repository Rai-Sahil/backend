package service

import (
	"errors"
	"fmt"

	"github.com/Rai-Sahil/backend/internal/domain"
	"github.com/Rai-Sahil/backend/internal/repository"
	"github.com/Rai-Sahil/backend/pkg/utils"
)

type AuthService struct {
	UserRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Register(user domain.User) error {
	existing, _ := s.UserRepo.FindByEmail(user.Email)
	if existing.ID != 0 {
		return errors.New("email already exists")
	}

	fmt.Print(user.Password)

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (domain.User, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return domain.User{}, err
	}

	if !utils.CheckPassword(password, user.Password) {
		return domain.User{}, errors.New("invalid credentials")
	}

	return user, nil
}
