package sectioncontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/wiratkhamphan/go_next_2024_api-master/models"
)

// UserController struct for handling user-related actions
type Sectionct struct {
	db *sqlx.DB
}

// NewUserController creates and returns a new instance of UserController
func NewSection(db *sqlx.DB) *Sectionct {
	return &Sectionct{db: db}
}

// List method to fetch sections by departmentId
func (s *Sectionct) List(c *fiber.Ctx) error {
	var sections []models.Section
	departmentID := c.Params("departmentId")

	// Query to fetch sections by department ID
	query := `SELECT id, name, departmentId, department, status, createdAt FROM Sections WHERE departmentId = ? ORDER BY name ASC`

	// Executing the query
	rows, err := s.db.Query(query, departmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch sections"})
	}
	defer rows.Close()

	// Iterating through the rows
	for rows.Next() {
		var section models.Section
		// Scanning the row into the section model
		if err := rows.Scan(&section.ID, &section.Name, &section.DepartmentId, &section.Department, &section.Status, &section.CreatedAt); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning section data"})
		}
		// Append to the sections slice
		sections = append(sections, section)
	}

	// Return the sections as JSON
	return c.JSON(fiber.Map{"sections": sections})
}
