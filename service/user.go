package service

import (
	"EventShot_Monitor/domain"
	"EventShot_Monitor/utils"
	"errors"
)

type UserService struct {
	userRepo domain.UserRepository
	roleRepo domain.RoleRepository
}

func NewUserService(userRepo domain.UserRepository, roleRepo domain.RoleRepository) UserService {
	return UserService{userRepo: userRepo, roleRepo: roleRepo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	if err := validateUserData(user); err != nil {
		return err
	}

	role, err := s.roleRepo.GetByID(user.Role.ID)
	if err != nil {
		return err
	}

	user.Role = *role

	hashedPassword, err := utils.HashPassword(user.HashedPassword)
	if err != nil {
		return err
	}

	user.HashedPassword = hashedPassword

	return s.userRepo.Create(user)
}

func validateUserData(user *domain.User) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.HashedPassword == "" {
		return errors.New("password is required")
	}
	if user.Role.Name == "" {
		return errors.New("role is required")
	}
	// todo роли должны браться из базы данных
	// Добавьте дополнительные проверки по мере необходимости.
	return nil
}
