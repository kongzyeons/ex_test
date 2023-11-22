package routes

import (
	"go_ex3/controller"
	"go_ex3/services"

	"github.com/gofiber/fiber/v2"
)

func BeefRouter(app *fiber.App) {
	beefSrv := services.NewBeefSrv()
	beefRest := controller.NewBeefRest(beefSrv)

	app.Get("/beef/summary", beefRest.GetGroupBeef)
	app.Post("/beef/summary", beefRest.GetGroupBeef)

}
