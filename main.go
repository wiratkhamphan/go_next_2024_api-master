package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/go_next_2024_api-master/database"
	"github.com/wiratkhamphan/go_next_2024_api-master/routes"
)

func main() {
	fmt.Println("Dev code app running...")

	database.DatabaseConfig()

	app := fiber.New()

	// Setup CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == "OPTIONS" {
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}

		return c.Next()
	})

	routes.Setup(app)

	app.Listen(":8000")
}
