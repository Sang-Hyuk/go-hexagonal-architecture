package repository

import (
	"time"

	"hexagonal-arch/example-1/internal/adapter/out/repository/entity"
	"hexagonal-arch/example-1/internal/core/domain"
	outport "hexagonal-arch/example-1/internal/core/port/out"
)

type User struct {
}

func NewUserRepository() outport.UserRepository {
	return &User{}
}

func (r *User) Get(id string) (domain.User, error) {
	// TODO: get user by id from database
	now := time.Now()

	entityUser := entity.User{
		ID:       id,
		Name:     "홍길동",
		Phone:    "010-1234-1234",
		Created:  now,
		Modified: now,
	}

	return entity.UserEntityToDomain(entityUser), nil
}

func (r *User) Insert(user domain.User) error {
	now := time.Now()

	entityUser := entity.UserDomainToEntity(user)

	entityUser.Created = now
	entityUser.Modified = now

	// TODO: insert entityUser to database

	return nil
}

func (r *User) Update(user domain.User) error {
	now := time.Now()

	entityUser := entity.UserDomainToEntity(user)

	entityUser.Modified = now

	// TODO: update entityUser to database

	return nil
}
