package in

import (
	"hexagonal-arch/example-1/internal/core/domain"
)

type UserUseCase interface {
	Get(id string) (domain.User, error)
	Upsert(user domain.User) error
}
