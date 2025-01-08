package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/wiratkhamphan/go_next_2024_api-master/models"
)

// UserController struct for handling user-related actions
type UserController struct {
	db *sqlx.DB
}

// NewUserController creates and returns a new instance of UserController
func NewUserController(db *sqlx.DB) *UserController {
	return &UserController{db: db}
}

// SignIn handles user login
func (u *UserController) SignIn(c *fiber.Ctx) error {
	data := new(models.UserLogin)

	// Parse the request body to extract login data
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.Users
	// Query to check if the user exists with the provided username and password
	query := "SELECT id, username, level FROM users WHERE username = ? AND password = ? AND status = 'active'"
	row := u.db.QueryRow(query, data.Username, data.Password)

	// Scan the result and handle errors
	if err := row.Scan(&user.Id, &user.Username, &user.Level); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(fiber.Map{
		"message": "Sign in successful",
		"user":    user,
	})
}

// Update allows updating user information
func (u *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(models.Users)

	// Parse the request body to extract user data
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Update the user details in the database
	query := "UPDATE users SET username = ?, password = ?, level = ?, section_id = ? WHERE id = ?"
	_, err := u.db.Exec(query, data.Username, data.Password, data.Level, data.SectionID, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

// Remove sets the user's status to 'inactive' to remove them
func (u *UserController) Remove(c *fiber.Ctx) error {
	id := c.Params("id")

	// Mark the user as inactive in the database
	query := "UPDATE users SET status = 'inactive' WHERE id = ?"
	_, err := u.db.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to remove user"})
	}

	return c.JSON(fiber.Map{"message": "User removed successfully"})
}

// ListEngineer retrieves all engineers who are active
func (u *UserController) ListEngineer(c *fiber.Ctx) error {
	var engineers []models.Users
	// Query to fetch engineers from the database
	query := "SELECT id, username, level FROM users WHERE level = 'engineer' AND status = 'active' ORDER BY username ASC"
	rows, err := u.db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	// Iterate over the result set and append engineers to the list
	for rows.Next() {
		var user models.Users
		if err := rows.Scan(&user.Id, &user.Username, &user.Level); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning user"})
		}
		engineers = append(engineers, user)
	}

	return c.JSON(engineers)
}

// List retrieves all users
func (u *UserController) List(c *fiber.Ctx) error {
	var users []models.Users
	// Query to fetch all users from the database
	query := "SELECT id, username, level FROM users WHERE status = 'active' ORDER BY username ASC"
	rows, err := u.db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	for rows.Next() {
		var user models.Users
		if err := rows.Scan(&user.Id, &user.Username, &user.Level); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning user"})
		}
		users = append(users, user)
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

// Create creates a new user
func (u *UserController) Create(c *fiber.Ctx) error {
	data := new(models.Users)

	// Parse the request body to extract user data
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Insert the new user into the database
	query := "INSERT INTO users (username, password, level, section_id, status) VALUES (?, ?, ?, ?, 'active')"
	_, err := u.db.Exec(query, data.Username, data.Password, data.Level, data.SectionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(fiber.Map{"message": "User created successfully"})
}

// UpdateUser updates a specific user by ID
func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(models.Users)

	// Parse the request body to extract user data
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Update the user details in the database
	query := "UPDATE users SET username = ?, password = ?, level = ?, section_id = ? WHERE id = ?"
	_, err := u.db.Exec(query, data.Username, data.Password, data.Level, data.SectionID, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}
