package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"

	"tugas_akhir_example/internal/pkg/repository"

	"tugas_akhir_example/internal/pkg/usecase"
)

// UserRoute routes the user group path
func UserRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewUserRepository(containerConf.Mysqldb)
	usecase := usecase.NewUserUseCase(repo)
	controller := controller.NewUserController(usecase)

	userAPI := r.Group("/user")
	userAPI.Get("", controller.GetMyProfile)
	userAPI.Put("", controller.UpdateProfile)
	userAPI.Get("alamat", controller.GetMyAlamats)
	userAPI.Get("alamat/:id", utils.AlamatAuthMiddleware(repo), controller.GetAlamatById)
	userAPI.Post("alamat", controller.CreateAlamat)
	userAPI.Put("alamat/:id", utils.AlamatAuthMiddleware(repo), controller.UpdateAlamatById)
	userAPI.Delete("alamat/:id", utils.AlamatAuthMiddleware(repo), controller.DeleteAlamatById)
}
