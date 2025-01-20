package devicecontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type DeviceController struct {
	db *sqlx.DB
}

func NewDevice(db *sqlx.DB) *DeviceController {
	return &DeviceController{db: db}
}
func (d *DeviceController) Create(c *fiber.Ctx) error {
	var body struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		DepartmentID int    `json:"departmentId"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	query := `
		INSERT INTO Devices (name, description, departmentId) 
		VALUES (?, ?, ?)
	`

	_, err := d.db.Exec(query, body.Name, body.Description, body.DepartmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create device"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Device created successfully"})
}
func (d *DeviceController) List(c *fiber.Ctx) error {
	query := `SELECT id, name, description, departmentId FROM Devices ORDER BY name ASC`

	rows, err := d.db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch devices"})
	}
	defer rows.Close()

	var devices []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		DepartmentID int    `json:"departmentId"`
	}

	for rows.Next() {
		var device struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			DepartmentID int    `json:"departmentId"`
		}
		if err := rows.Scan(&device.ID, &device.Name, &device.Description, &device.DepartmentID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning device data"})
		}
		devices = append(devices, device)
	}

	return c.JSON(fiber.Map{"devices": devices})
}
func (d *DeviceController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var body struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		DepartmentID int    `json:"departmentId"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	query := `
		UPDATE Devices 
		SET name = ?, description = ?, departmentId = ? 
		WHERE id = ?
	`

	_, err := d.db.Exec(query, body.Name, body.Description, body.DepartmentID, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update device"})
	}

	return c.JSON(fiber.Map{"message": "Device updated successfully"})
}
func (d *DeviceController) Remove(c *fiber.Ctx) error {
	id := c.Params("id")

	query := `DELETE FROM Devices WHERE id = ?`

	_, err := d.db.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete device"})
	}

	return c.JSON(fiber.Map{"message": "Device deleted successfully"})
}
