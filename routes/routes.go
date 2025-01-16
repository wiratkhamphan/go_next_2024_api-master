package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	Company "github.com/wiratkhamphan/go_next_2024_api-master/controllers/Company"
	Record "github.com/wiratkhamphan/go_next_2024_api-master/controllers/Record"
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
}

// SetupRoutes defines user-specific routes
func SetupRoutes_user(app *fiber.App, db *sqlx.DB) {
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
	userGroup.Get("/api/user/level", userController.Level)
}

func SerupRoutes_Record(app *fiber.App, db *sqlx.DB) {

	RepairRecordController := Record.NewRecord(db)

	RecordGroup := app.Group("/api/Record")

	RecordGroup.Get("repairRecord/list", RepairRecordController.List)
	RecordGroup.Post("repairRecord/create", RepairRecordController.Create)

}

func SerupRoutes_Company(app *fiber.App, db *sqlx.DB) {

	CompanyController := Company.NewCompany(db)

	CompanyGroup := app.Group("/api/company")

	CompanyGroup.Get("Company/list", CompanyController.CheckSignIn, CompanyController.GetCompanyInfo)
	CompanyGroup.Post("Company/create", CompanyController.CheckSignIn, CompanyController.UpdateCompanyInfo)

}
