package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"

	"tugas_akhir_example/internal/pkg/repository"

	"tugas_akhir_example/internal/pkg/usecase"
)

// ProdukRoute routes the produk group path
func ProdukRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewProdukRepository(containerConf.Mysqldb)
	usecase := usecase.NewProdukUseCase(repo)
	controller := controller.NewProdukController(usecase)

	produkAPI := r.Group("/product")
	produkAPI.Get("", controller.GetAllProduks)
	produkAPI.Get(":id", controller.GetProdukById)
	produkAPI.Post("", controller.CreateProduk)
	produkAPI.Put(":id", utils.ProdukAuthMiddleware(repo), controller.UpdateProdukById)
	produkAPI.Delete(":id", utils.ProdukAuthMiddleware(repo), controller.DeleteProdukById)
}
