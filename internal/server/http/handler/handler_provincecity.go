package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"

	"tugas_akhir_example/internal/pkg/repository"

	"tugas_akhir_example/internal/pkg/usecase"
)

// ProvinceCityRoute routes the provincecity group path
func ProvinceCityRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewProvinceCityRepository()
	usecase := usecase.NewProvinceCityUseCase(repo)
	controller := controller.NewProvinceCityController(usecase)

	provinceCityAPI := r.Group("/provcity")
	provinceCityAPI.Get("listprovincies", controller.GetAllProvinces)
	provinceCityAPI.Get("listcities/:prov_id", controller.GetAllCities)
	provinceCityAPI.Get("detailprovince/:prov_id", controller.GetProvinceById)
	provinceCityAPI.Get("detailcity/:city_id", controller.GetCityById)
}
