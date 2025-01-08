package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	controllers "github.com/wiratkhamphan/go_next_2024_api-master/controllers/UserController"
)

// Setup defines initial routes including user authentication
func Setup(app *fiber.App) {
	// Hello World route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Define routes for registration and login (controllers should be properly implemented)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//
	// Repair record routes (Uncomment and implement as needed)
	//
	// app.Get("/api/repairRecord/list", RepairRecordController.list)
	// app.Post("/api/repairRecord/create", RepairRecordController.create)
	// app.Put("/api/repairRecord/update/:id", RepairRecordController.update)
	// app.Delete("/api/repairRecord/remove/:id", RepairRecordController.remove)
	// app.Put("/api/repairRecord/updateStatus/:id", RepairRecordController.upateStatus)
	// app.Put("/api/repairRecord/receive", RepairRecordController.receive)

	//
	// Department and section routes (Uncomment and implement as needed)
	//
	// app.Get("/api/department/list", DepartmentController.list)
	// app.Get("/api/section/listByDepartment/:departmentId", SectionController.listByDepartment)

	//
	// Device routes (Uncomment and implement as needed)
	//
	// app.Post("/api/device/create", DeviceController.create)
	// app.Get("/api/device/list", DeviceController.list)
	// app.Put("/api/device/update/:id", DeviceController.update)
	// app.Delete("/api/device/remove/:id", DeviceController.remove)
}

// SetupRoutes defines user-specific routes
func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	// Initialize user controller with the database connection
	userController := controllers.NewUserController(db)

	// Define routes for the user group
	userGroup := app.Group("/api/user")
	// Map the routes to the controller methods
	userGroup.Post("/signin", userController.SignIn)            // login
	userGroup.Put("/update", userController.Update)             // update
	userGroup.Get("/list", userController.List)                 // list users
	userGroup.Post("/create", userController.Create)            // create new user
	userGroup.Put("/updateUser/:id", userController.UpdateUser) // update specific user by ID
	userGroup.Delete("/remove/:id", userController.Remove)      // remove a user
	userGroup.Get("/listEngineer", userController.ListEngineer) // list engineers
}
