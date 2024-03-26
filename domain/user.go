package domain

type User struct {
	ID             uint   `gorm:"primaryKey;autoincrement;not null"`
	HashedPassword string `gorm:"not null"`
	Role           Role   `gorm:"not null;foreignKey:RoleID"`
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

type UserService interface {
	GetUser(id uint) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}
