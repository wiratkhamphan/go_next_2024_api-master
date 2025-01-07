package controllers

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wiratkhamphan/go_next_2024_api-master/database"
	"github.com/wiratkhamphan/go_next_2024_api-master/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

// Register a new user
func Register(c *fiber.Ctx) error {
	data := models.User{}

	// Parse JSON body into the User struct
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	// Validate input
	if data.Email == "" || data.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	// Hash the password
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
	}

	// Using Raw SQL for clarity
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err = database.DB.Exec(query, data.Name, data.Email, string(password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
	}

	// Return the newly created user
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(password),
	}

	return c.JSON(fiber.Map{
		"message": "user registered successfully",
		"user":    user,
	})
}

// Login user and issue a JWT
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	var user models.User

	// Using Raw SQL to fetch the user
	query := "SELECT id, name, email, password FROM users WHERE email = ? LIMIT 1"
	row := database.DB.QueryRow(query, data["email"])

	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "database error"})
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "incorrect password"})
	}

	// Create JWT claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not generate token"})
	}

	// Create JWT cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,     // Ensure cookies are sent over HTTPS in production
		SameSite: "Strict", // Restrict cookie to same-site requests for security
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "login successful"})
}

// Get user information from JWT
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthenticated"})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	// Using Raw SQL to fetch user by ID
	query := "SELECT id, name, email, password FROM users WHERE id = ? LIMIT 1"
	row := database.DB.QueryRow(query, claims.Issuer)

	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}

// Logout user by clearing the JWT cookie
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "logout successful"})
}
