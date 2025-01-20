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
	id := c.Params("id") // รับค่า ID จาก URL parameter
	var body models.RepairRecord

	// Parse JSON body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// SQL Update query
	query := `
		UPDATE RepairRecord 
		SET 
			customerName = :customerName,
			customerPhone = :customerPhone,
			deviceName = :deviceName,
			deviceBarcode = :deviceBarcode,
			deviceSerial = :deviceSerial,
			problem = :problem,
			solving = :solving,
			deviceId = :deviceId,
			userId = :userId,
			engineerId = :engineerId,
			status = :status
		WHERE id = :id
	`

	params := map[string]interface{}{
		"id":            id,
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

	// Execute query
	_, err := r.db.NamedExec(query, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update repair record",
			"details": err.Error(),
		})
	}

	// Return success
	return c.JSON(fiber.Map{
		"message": "Repair record updated successfully",
	})
}

func (r *Record) Remove(c *fiber.Ctx) error {
	id := c.Params("id") // รับค่า ID จาก URL parameter

	query := `DELETE FROM RepairRecord WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete repair record",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Repair record deleted successfully",
	})
}

func (r *Record) UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id") // รับค่า ID จาก URL parameter
	var body struct {
		Status string `json:"status"` // รับสถานะใหม่จาก request body
	}

	// Parse JSON body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// ตรวจสอบว่า Status ไม่เป็นค่าว่าง
	if body.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Status is required",
		})
	}

	// SQL Update query
	query := `UPDATE RepairRecord SET status = ? WHERE id = ?`

	// Execute query
	_, err := r.db.Exec(query, body.Status, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update status",
			"details": err.Error(),
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"message": "Status updated successfully",
	})
}

func (r *Record) Receive(c *fiber.Ctx) error {
	var body struct {
		ID     int    `json:"id"`     // ID ของ repair record
		Status string `json:"status"` // สถานะใหม่ที่ต้องการตั้งค่า
	}

	// Parse JSON body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// ตรวจสอบว่า ID และ Status มีค่าหรือไม่
	if body.ID == 0 || body.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID and Status are required",
		})
	}

	// SQL Update query
	query := `UPDATE RepairRecord SET status = ? WHERE id = ?`

	// Execute query
	_, err := r.db.Exec(query, body.Status, body.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update record",
			"details": err.Error(),
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"message": "Record status updated successfully",
	})
}
func (r *Record) IncomePerMonth(c *fiber.Ctx) error {
	query := `
		SELECT 
			MONTH(payDate) AS month, 
			YEAR(payDate) AS year, 
			SUM(amount) AS totalIncome
		FROM 
			RepairRecord
		WHERE 
			payDate IS NOT NULL
		GROUP BY 
			YEAR(payDate), MONTH(payDate)
		ORDER BY 
			YEAR(payDate), MONTH(payDate)
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch income per month"})
	}
	defer rows.Close()

	var result []struct {
		Month       int     `json:"month"`
		Year        int     `json:"year"`
		TotalIncome float64 `json:"totalIncome"`
	}

	for rows.Next() {
		var row struct {
			Month       int     `json:"month"`
			Year        int     `json:"year"`
			TotalIncome float64 `json:"totalIncome"`
		}
		if err := rows.Scan(&row.Month, &row.Year, &row.TotalIncome); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning income data"})
		}
		result = append(result, row)
	}

	return c.JSON(fiber.Map{"incomePerMonth": result})
}
func (r *Record) Report(c *fiber.Ctx) error {
	startDate := c.Params("startDate")
	endDate := c.Params("endDate")

	query := `
		SELECT 
			id, customerName, amount, payDate 
		FROM 
			RepairRecord
		WHERE 
			payDate BETWEEN ? AND ?
		ORDER BY 
			payDate DESC
	`

	rows, err := r.db.Query(query, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch report"})
	}
	defer rows.Close()

	var report []struct {
		ID           int     `json:"id"`
		CustomerName string  `json:"customerName"`
		Amount       float64 `json:"amount"`
		PayDate      string  `json:"payDate"`
	}

	for rows.Next() {
		var record struct {
			ID           int     `json:"id"`
			CustomerName string  `json:"customerName"`
			Amount       float64 `json:"amount"`
			PayDate      string  `json:"payDate"`
		}
		if err := rows.Scan(&record.ID, &record.CustomerName, &record.Amount, &record.PayDate); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning report data"})
		}
		report = append(report, record)
	}

	return c.JSON(fiber.Map{"report": report})
}
func (r *Record) Dashboard(c *fiber.Ctx) error {
	query := `
		SELECT 
			(SELECT COUNT(*) FROM RepairRecord WHERE status = 'In Progress') AS inProgress,
			(SELECT COUNT(*) FROM RepairRecord WHERE status = 'Completed') AS completed,
			(SELECT SUM(amount) FROM RepairRecord WHERE status = 'Completed') AS totalIncome
	`

	row := r.db.QueryRow(query)

	var dashboard struct {
		InProgress  int     `json:"inProgress"`
		Completed   int     `json:"completed"`
		TotalIncome float64 `json:"totalIncome"`
	}

	if err := row.Scan(&dashboard.InProgress, &dashboard.Completed, &dashboard.TotalIncome); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch dashboard data"})
	}

	return c.JSON(dashboard)
}
