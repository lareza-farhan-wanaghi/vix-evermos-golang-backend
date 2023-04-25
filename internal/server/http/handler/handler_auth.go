package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"

	"tugas_akhir_example/internal/pkg/repository"

	"tugas_akhir_example/internal/pkg/usecase"
)

// AuthRoute routes the auth group path
func AuthRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewAuthRepository(containerConf.Mysqldb)
	usecase := usecase.NewAuthUseCase(repo, containerConf.Apps.SecretJwt)
	controller := controller.NewAuthController(usecase)

	authAPI := r.Group("/auth")
	authAPI.Post("register", controller.RegisterUsers)
	authAPI.Post("login", controller.LoginUsers)
}
