package service

import (
	"EventShot_Monitor/domain"
	"errors"
)

type RoleService struct {
	roleRepo domain.RoleRepository
}

func NewRoleService(roleRepo domain.RoleRepository) RoleService {
	return RoleService{roleRepo: roleRepo}
}

func (s *RoleService) CreateRole(role *domain.Role) error {
	if err := validateRoleData(role); err != nil {
		return err
	}

	return s.roleRepo.Create(role)
}

func (s *RoleService) DeleteRole(id uint) error {
	_, err := s.roleRepo.GetByID(id)
	if err != nil {
		return errors.New("role is not existing")
	}

	return s.roleRepo.Delete(id)
}

func (s *RoleService) GetRoles() ([]*domain.Role, error) {
	roles, err := s.roleRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func validateRoleData(role *domain.Role) error {
	if role.Name == "" {
		return errors.New("role name is required")
	}
	return nil
}
