package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	CompanyController "github.com/wiratkhamphan/go_next_2024_api-master/controllers/CompanyController"
	Department "github.com/wiratkhamphan/go_next_2024_api-master/controllers/DepartmentController"
	Device "github.com/wiratkhamphan/go_next_2024_api-master/controllers/DeviceController"
	RepairRecordController "github.com/wiratkhamphan/go_next_2024_api-master/controllers/RepairRecordController"
	Section "github.com/wiratkhamphan/go_next_2024_api-master/controllers/SectionController"

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

// company
func SerupRoutes_Company(app *fiber.App, db *sqlx.DB) {

	CompanyController := CompanyController.NewCompany(db)

	CompanyGroup := app.Group("/api/company")

	CompanyGroup.Get("/info", controllers.CheckSignIn, CompanyController.GetCompanyInfo)
	CompanyGroup.Post("/update", CompanyController.UpdateCompanyInfo)

}

// repair record
func SerupRoutes_Record(app *fiber.App, db *sqlx.DB) {

	RepairRecordController := RepairRecordController.NewRecord(db)

	RecordGroup := app.Group("/api/repairRecord")

	RecordGroup.Get("/list", RepairRecordController.List)
	RecordGroup.Post("/create", RepairRecordController.Create)
	RecordGroup.Put("/update/:id", RepairRecordController.Update)
	RecordGroup.Delete("/remove/:id", RepairRecordController.Remove)
	RecordGroup.Put("/updateStatus/:id", RepairRecordController.UpdateStatus)
	RecordGroup.Put("/receive", RepairRecordController.Receive)
	RecordGroup.Get("/incomePerMonth", RepairRecordController.IncomePerMonth)
	app.Get("/api/income/report/:startDate/:endDate", RepairRecordController.Report)
	RecordGroup.Get("/dashboard", RepairRecordController.Dashboard)

}

func SetupRoutes_DepartmentController(app *fiber.App, db *sqlx.DB) {

	DepartmentController := Department.NewDepartment(db)

	Group_DepartmentController := app.Group("/api/department")

	Group_DepartmentController.Get("/list", DepartmentController.List)
}
func SetupRoutes_SectionController(app *fiber.App, db *sqlx.DB) {
	SectionController := Section.NewSection(db)

	Group_SectionController := app.Group("/api/section")
	Group_SectionController.Get("/listByDepartment/:departmentId", SectionController.List)
}

func SetupRoutes_DeviceController(app *fiber.App, db *sqlx.DB) {

	DeviceController := Device.NewDevice(db)

	Group_DeviceController := app.Group("/api/device")

	Group_DeviceController.Post("/create", DeviceController.Create)
	Group_DeviceController.Get("/list", DeviceController.List)
	Group_DeviceController.Put("/update/:id", DeviceController.Update)
	Group_DeviceController.Delete("/remove/:id", DeviceController.Remove)
}
