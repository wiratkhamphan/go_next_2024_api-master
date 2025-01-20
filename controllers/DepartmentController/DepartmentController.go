package departmentcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type DepartmentController struct {
	db *sqlx.DB
}

func NewDepartment(db *sqlx.DB) *DepartmentController {
	return &DepartmentController{db: db}
}
func (d *DepartmentController) List(c *fiber.Ctx) error {
	query := `SELECT id, name FROM Departments ORDER BY name ASC`

	rows, err := d.db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch departments"})
	}
	defer rows.Close()

	var departments []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	for rows.Next() {
		var department struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		if err := rows.Scan(&department.ID, &department.Name); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning department data"})
		}
		departments = append(departments, department)
	}

	return c.JSON(fiber.Map{"departments": departments})
}
