package out

import (
	"hexagonal-arch/example-1/internal/core/domain"
)

type UserRepository interface {
	Get(id string) (domain.User, error)
	Insert(user domain.User) error
	Update(user domain.User) error
}
