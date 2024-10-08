package usecase

import (
	"hexagonal-arch/example-1/internal/core/domain"
	inport "hexagonal-arch/example-1/internal/core/port/in"
	outport "hexagonal-arch/example-1/internal/core/port/out"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	userRepository outport.UserRepository
}

func NewUserUseCase(userRepository outport.UserRepository) inport.UserUseCase {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Get(id string) (domain.User, error) {
	user, err := u.userRepository.Get(id)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "failed to get user by id")
	}

	return user, nil
}

func (u *User) Upsert(user domain.User) error {
	if _, err := u.userRepository.Get(user.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := u.userRepository.Insert(user); err != nil {
				return errors.Wrap(err, "failed to insert user")
			}
			return nil
		}

		return errors.Wrap(err, "failed to get user by id")
	}

	if err := u.userRepository.Update(user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
