package in

import (
	"hexagonal-arch/internal/core/domain"
)

type UserUseCase interface {
	Get(id string) (domain.User, error)
	Upsert(user domain.User) error
}
