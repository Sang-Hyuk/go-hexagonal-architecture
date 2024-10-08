package entity

import (
	"time"

	"hexagonal-arch/example-1/internal/core/domain"
)

type User struct {
	ID       string
	Name     string
	Phone    string
	Created  time.Time
	Modified time.Time
}

func UserDomainToEntity(user domain.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}
}

func UserEntityToDomain(user User) domain.User {
	return domain.User{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:    u.ID,
		Name:  u.Name,
		Phone: u.Phone,
	}
}
