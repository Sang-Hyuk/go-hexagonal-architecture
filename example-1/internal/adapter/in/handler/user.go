package handler

import (
	"hexagonal-arch/example-1/internal/adapter/in/handler/request"
	"hexagonal-arch/example-1/internal/adapter/in/handler/response"
	"hexagonal-arch/example-1/internal/core/domain"
	inport "hexagonal-arch/example-1/internal/core/port/in"

	"github.com/gin-gonic/gin"
)

type User struct {
	userUseCase inport.UserUseCase
}

func NewUserHandler(userUseCase inport.UserUseCase) User {
	return User{
		userUseCase: userUseCase,
	}
}

func (h *User) Get(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userUseCase.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp := response.UserDomainToResponse(user)

	c.JSON(200, gin.H{"data": resp})
}

func (h *User) Update(c *gin.Context) {
	var user request.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.ID == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	if user.Name == "" {
		c.JSON(400, gin.H{"error": "name is required"})
		return
	}

	domainUser := domain.User{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}

	err := h.userUseCase.Upsert(domainUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user updated"})
}
