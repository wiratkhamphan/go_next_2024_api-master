package controllers

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/wiratkhamphan/go_next_2024_api-master/models"
	"golang.org/x/crypto/bcrypt"
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

	// Parse the request body
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.Users

	// Query user data by username
	query := "SELECT id, username, level, password FROM users WHERE username = ? AND status = 'active'"
	row := u.db.QueryRow(query, data.Username)

	// Scan result into user struct
	if err := row.Scan(&user.ID, &user.Username, &user.Level, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Compare password with hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	// Create JWT cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	}
	c.Cookie(&cookie)

	// Return response with token and user info
	return c.JSON(fiber.Map{
		"message": "Sign in successful",
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"level":    user.Level,
		},
		"token": token,
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
		if err := rows.Scan(&user.ID, &user.Username, &user.Level); err != nil {
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
		if err := rows.Scan(&user.ID, &user.Username, &user.Level); err != nil {
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
	if data.Username == "" || data.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	// Hash the password
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
	}
	// Insert the new user into the database
	query := "INSERT INTO users (username, password, level, sectionId , status) VALUES (?, ?, ?, ?, 'active')"
	_, err = u.db.Exec(query, data.Username, string(password), data.Level, data.SectionID)
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

func (u *UserController) Level(c *fiber.Ctx) error {
	type Claims struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}

	// รับ Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	// ดึง Token จาก Header
	tokenString := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	// ตรวจสอบและแปลง Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Replace with your JWT secret key
		return []byte("your_secret_key"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// ดึง User ID จาก Claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to parse token claims",
		})
	}

	userID := claims.ID

	// ค้นหา Level ของผู้ใช้จากฐานข้อมูล
	userLevel, err := u.queryUserLevelByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query user level",
		})
	}

	// ส่ง Level กลับใน JSON
	return c.JSON(fiber.Map{
		"user_level": *userLevel,
	})
}

// ฟังก์ชันสำหรับดึงข้อมูล Level ของผู้ใช้จากฐานข้อมูล
func (u *UserController) queryUserLevelByID(userID string) (*int, error) {
	var level int

	// คำสั่ง SQL เพื่อดึงข้อมูล Level
	query := `
		SELECT level 
		FROM users 
		WHERE id = ?
	`

	// เตรียมการ query
	err := u.db.QueryRow(query, userID).Scan(&level)
	if err != nil {
		return nil, err
	}

	return &level, nil
}
