package models

// User represents a user object for database interaction
type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
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

// Database models
type RepairRecord struct {
	ID                int     `json:"id"`
	CustomerName      string  `json:"customerName"`
	CustomerPhone     string  `json:"customerPhone"`
	DeviceName        string  `json:"deviceName"`
	DeviceBarcode     string  `json:"deviceBarcode"`
	DeviceSerial      *string `json:"deviceSerial"`
	Problem           string  `json:"problem"`
	Solving           *string `json:"solving"`
	DeviceID          *uint   `json:"deviceId"`
	UserID            *uint   `json:"userId"`
	EngineerID        *uint   `json:"engineerId"`
	Status            string  `json:"status" gorm:"default:active"`
	CreatedAt         *string `json:"createdAt"`
	EndJobDate        *string `json:"endJobDate"`
	PayDate           *string `json:"payDate"`
	Amount            *int    `json:"amount"`
	ImageBeforeRepair *string `json:"imageBeforeRepair"`
	ImageAfterRepair  *string `json:"imageAfterRepair"`
	Engineer          *User   `json:"engineer,omitempty" gorm:"-"`
}

// Company struct

type Company struct {
	ID           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Address      string `db:"address" json:"address"`
	Phone        string `db:"phone" json:"phone"`
	Email        string `db:"email" json:"email"`
	FacebookPage string `db:"facebook_page" json:"facebook_page"`
	TaxCode      string `db:"tax_code" json:"tax_code"`
}
