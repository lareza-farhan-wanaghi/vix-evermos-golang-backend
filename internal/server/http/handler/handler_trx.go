package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"

	"tugas_akhir_example/internal/pkg/repository"

	"tugas_akhir_example/internal/pkg/usecase"
)

// TrxRoute routes the trx group path
func TrxRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewTrxRepository(containerConf.Mysqldb)
	usecase := usecase.NewTrxUseCase(repo)
	controller := controller.NewTrxController(usecase)

	trxAPI := r.Group("/trx")
	trxAPI.Get("", controller.GetAllTrxs)
	trxAPI.Get(":id", controller.GetTrxById)
	trxAPI.Post("", controller.CreateTrx)
}
