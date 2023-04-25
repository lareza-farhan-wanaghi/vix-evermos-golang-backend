package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	bookcontroller "tugas_akhir_example/internal/pkg/controller"

	bookrepository "tugas_akhir_example/internal/pkg/repository"

	bookusecase "tugas_akhir_example/internal/pkg/usecase"
)

// BookRoute routes the book group path
func BookRoute(r fiber.Router, containerConf *container.Container) {
	repo := bookrepository.NewBookRepository(containerConf.Mysqldb)
	usecase := bookusecase.NewBookUseCase(repo)
	controller := bookcontroller.NewBookController(usecase)

	bookAPI := r.Group("/book")
	bookAPI.Get("", controller.GetAllBook)
	bookAPI.Get("/:id_book", controller.GetBookByID)
	bookAPI.Post("", controller.CreateBook)
	bookAPI.Put("/:id_book", controller.UpdateBookByID)
	bookAPI.Delete("/:id_book", controller.DeleteBookByID)
}
