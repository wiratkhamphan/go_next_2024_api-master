package models

// User represents a user object for database interaction
type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users represents a user for the application
type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Level     string `json:"level"`
	SectionID int    `json:"sectionId"`
}

// UserLogin represents the login credentials provided by the user
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type CustomerRepository interface {
	GetAll() ([]Users, error)
	GetById(int) (*Users, error)
	Save(Users) (*Users, error)
}
