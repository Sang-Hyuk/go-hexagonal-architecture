package main

import (
	"hexagonal-arch/example-1/internal/adapter/in/handler"
	"hexagonal-arch/example-1/internal/adapter/out/repository"
	"hexagonal-arch/example-1/internal/core/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepository := repository.NewUserRepository()

	userUseCase := usecase.NewUserUseCase(userRepository)

	userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()
	r.GET("/user/:id", userHandler.Get)
	r.POST("/user", userHandler.Update)
	r.Run() // listen and serve on 0.0.0.0:8080
}
