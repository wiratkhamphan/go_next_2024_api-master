package company

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/wiratkhamphan/go_next_2024_api-master/models"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Company struct {
	db *sqlx.DB
}

func NewCompany(db *sqlx.DB) *Company {
	return &Company{db: db}
}

// Middleware for JWT Authentication
func (r *Company) CheckSignIn(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")[7:] // Extract Bearer token
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized: Token missing"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized: Invalid token"})
	}

	return c.Next()
}

// Handlers
func (r *Company) GetCompanyInfo(c *fiber.Ctx) error {
	var company models.Company
	query := "SELECT * FROM companies LIMIT 1"

	if err := r.db.Get(&company, query); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Company not found"})
	}

	return c.JSON(company)
}

func (r *Company) UpdateCompanyInfo(c *fiber.Ctx) error {
	var body models.Company
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	var existingCompany models.Company
	query := "SELECT * FROM companies LIMIT 1"

	// Check if company exists
	if err := r.db.Get(&existingCompany, query); err != nil {
		// Create new company if not found
		insertQuery := `
			INSERT INTO companies (name, address, phone, email, facebook_page, tax_code) 
			VALUES (:name, :address, :phone, :email, :facebook_page, :tax_code)
		`
		if _, err := r.db.NamedExec(insertQuery, &body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create company"})
		}
		return c.JSON(fiber.Map{"message": "Company created successfully"})
	}

	// Update existing company
	updateQuery := `
		UPDATE companies 
		SET name = :name, address = :address, phone = :phone, email = :email, facebook_page = :facebook_page, tax_code = :tax_code
		WHERE id = :id
	`
	body.ID = existingCompany.ID // Ensure the correct ID is used for update
	if _, err := r.db.NamedExec(updateQuery, &body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update company"})
	}

	return c.JSON(fiber.Map{"message": "Company updated successfully"})
}
