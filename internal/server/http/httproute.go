package http

import (
	route "tugas_akhir_example/internal/server/http/handler"

	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

// HTTPRouteInit initializes the routing table of the app
func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	route.BookRoute(api, containerConf)
	route.AuthRoute(api, containerConf)
	route.TokoRoute(api, containerConf)
	route.ProvinceCityRoute(api, containerConf)
	route.ProdukRoute(api, containerConf)
	route.UserRoute(api, containerConf)
	route.CategoryRoute(api, containerConf)
	route.TrxRoute(api, containerConf)

	r.Static("/static", "./static")
}
