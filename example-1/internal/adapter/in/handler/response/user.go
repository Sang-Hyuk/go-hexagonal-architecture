package response

import "hexagonal-arch/example-1/internal/core/domain"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func UserDomainToResponse(user domain.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}
}
