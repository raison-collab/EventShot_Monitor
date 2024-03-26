package domain

type Role struct {
	ID   uint   `gorm:"primaryKey;not null;autoincrement"`
	Name string `gorm:"not null"`
}

type RoleRepository interface {
	GetAll() ([]*Role, error)
	GetByID(id uint) (*Role, error)
	Create(role *Role) error
	Delete(id uint) error
}

type RoleService interface {
	GetRole(id uint) (*Role, error)
	GetRoles() ([]*Role, error)
	CreateRole(role *Role) error
	DeleteRole(id uint) error
}
