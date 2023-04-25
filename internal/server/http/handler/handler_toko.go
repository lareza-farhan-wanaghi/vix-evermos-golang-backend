package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"

	"tugas_akhir_example/internal/pkg/repository"

	"tugas_akhir_example/internal/pkg/usecase"
)

// TokoRoute routes the toko group path
func TokoRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewTokoRepository(containerConf.Mysqldb)
	usecase := usecase.NewTokoUseCase(repo, containerConf.Apps.SecretJwt)
	controller := controller.NewTokoController(usecase)

	tokoAPI := r.Group("/toko")
	tokoAPI.Get("", controller.GetAllToko)
	tokoAPI.Get("my", controller.GetMyToko)
	tokoAPI.Get(":id_toko", controller.GetTokoById)
	tokoAPI.Put(":id_toko", utils.TokoAuthMiddleware(repo), controller.UpdateTokoByID)
}
