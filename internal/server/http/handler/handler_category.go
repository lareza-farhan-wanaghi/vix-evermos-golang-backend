package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"

	"tugas_akhir_example/internal/pkg/repository"

	"tugas_akhir_example/internal/pkg/usecase"
)

// CategoryRoute routes the category group path
func CategoryRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewCategoryRepository(containerConf.Mysqldb)
	usecase := usecase.NewCategoryUseCase(repo)
	controller := controller.NewCategoryController(usecase)

	categoryAPI := r.Group("/category")
	categoryAPI.Get("", controller.GetAllCategories)
	categoryAPI.Get(":id", controller.GetCategoryById)
	categoryAPI.Post("", utils.CategoryAuthMiddleware(repo), controller.CreateCategory)
	categoryAPI.Put(":id", utils.CategoryAuthMiddleware(repo), controller.UpdateCategoryById)
	categoryAPI.Delete(":id", utils.CategoryAuthMiddleware(repo), controller.DeleteCategoryById)
}
