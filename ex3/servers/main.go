package main

import (
	"go_ex3/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// text := "...   Fatback t-bone t-bone, pastrami  ..  ... .t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."
	server := ":8001"
	app := fiber.New()
	routes.BeefRouter(app)
	app.Listen(server)

}
