package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/go_next_2024_api-master/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
