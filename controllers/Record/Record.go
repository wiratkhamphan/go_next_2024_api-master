package record

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/wiratkhamphan/go_next_2024_api-master/models"
)

type Record struct {
	db *sqlx.DB
}

func NewRecord(db *sqlx.DB) *Record {
	return &Record{db: db}
}
func (r *Record) List(c *fiber.Ctx) error {
	var repairRecords []models.RepairRecord

	query := `
		SELECT 
			r.id,
			r.customerName,
			r.customerPhone,
			r.deviceName,
			r.deviceBarcode,
			r.deviceSerial,
			r.problem,
			r.solving,
			r.deviceId,
			r.userId,
			r.engineerId,
			r.status,
			r.createdAt,
			r.endJobDate,
			r.payDate,
			r.amount,
			r.imageBeforeRepair,
			r.imageAfterRepair,
			u.username AS engineer
		FROM 
			repairRecord r
		LEFT JOIN 
			users u ON r.engineerId = u.id
		ORDER BY 
			r.id DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	for rows.Next() {
		var repairRecord models.RepairRecord
		if err := rows.Scan(
			&repairRecord.ID,
			&repairRecord.CustomerName,
			&repairRecord.CustomerPhone,
			&repairRecord.DeviceName,
			&repairRecord.DeviceBarcode,
			&repairRecord.DeviceSerial,
			&repairRecord.Problem,
			&repairRecord.Solving,
			&repairRecord.DeviceID,
			&repairRecord.UserID,
			&repairRecord.EngineerID,
			&repairRecord.Status,
			&repairRecord.CreatedAt,
			&repairRecord.EndJobDate,
			&repairRecord.PayDate,
			&repairRecord.Amount,
			&repairRecord.ImageBeforeRepair,
			&repairRecord.ImageAfterRepair,
			&repairRecord.Engineer, // Assuming RepairRecord has an "Engineer" field
		); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning repair record"})
		}
		repairRecords = append(repairRecords, repairRecord)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error iterating rows"})
	}

	return c.JSON(fiber.Map{
		"repairRecords": repairRecords,
	})
}

// Create creates a new repair record
func (r *Record) Create(c *fiber.Ctx) error {
	var body models.RepairRecord

	// Parse JSON request body into body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Log parsed body for debugging
	fmt.Printf("Parsed Body: %+v\n", body)

	// Validate required fields (optional)
	if body.CustomerName == "" || body.CustomerPhone == "" || body.DeviceName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Required fields are missing: customerName, customerPhone, deviceName",
		})
	}

	// SQL Insert query using NamedExec
	query := `
		INSERT INTO RepairRecord (
			customerName, customerPhone, deviceName, deviceBarcode, deviceSerial, problem, solving, deviceId, userId, engineerId, status, createdAt
		) VALUES (
			:customerName, :customerPhone, :deviceName, :deviceBarcode, :deviceSerial, :problem, :solving, :deviceId, :userId, :engineerId, :status, NOW()
		)
	`

	// Log query execution
	fmt.Println("Executing query...")

	// Create a map to hold the struct values explicitly as named parameters
	params := map[string]interface{}{
		"customerName":  body.CustomerName,
		"customerPhone": body.CustomerPhone,
		"deviceName":    body.DeviceName,
		"deviceBarcode": body.DeviceBarcode,
		"deviceSerial":  body.DeviceSerial,
		"problem":       body.Problem,
		"solving":       body.Solving,
		"deviceId":      body.DeviceID,
		"userId":        body.UserID,
		"engineerId":    body.EngineerID,
		"status":        body.Status,
	}

	// Execute the query with NamedExec and map parameters manually
	_, err := r.db.NamedExec(query, params)
	if err != nil {
		// Log the database error for debugging
		fmt.Printf("Database Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create repair record",
			"details": err.Error(), // Log the exact error for debugging
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Repair record created successfully",
	})
}

func (r *Record) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(models.RepairRecord)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	query := ""
	_, err := r.db.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"Uddate":  data,
	})
}
func (r *Record) Remove(c *fiber.Ctx) error {
	return c.JSON(c)
}
func (r *Record) UpdateStatus(c *fiber.Ctx) error {
	return c.JSON(c)
}
func (r *Record) Receive(c *fiber.Ctx) error {
	return c.JSON(c)
}
