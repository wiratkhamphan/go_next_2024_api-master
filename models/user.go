package models

import "time"

// User represents a user object for database interaction
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Users represents a user for the application
type Users struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Level     string `json:"level"` // e.g., "user", "admin", "engineer"
	SectionID int    `json:"sectionId"`
}

// UserLogin represents the login credentials provided by the user
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RepairRecord represents a record of a device repair in the system
type RepairRecord struct {
	ID                int        `db:"id" json:"id"`
	CustomerName      string     `db:"customerName" json:"customerName"`
	CustomerPhone     string     `db:"customerPhone" json:"customerPhone"`
	DeviceName        string     `db:"deviceName" json:"deviceName"`
	DeviceBarcode     string     `db:"deviceBarcode" json:"deviceBarcode"`
	DeviceSerial      string     `db:"deviceSerial" json:"deviceSerial"`
	Problem           string     `db:"problem" json:"problem"`
	Solving           string     `db:"solving" json:"solving"`
	DeviceID          int        `db:"deviceId" json:"deviceId"`
	UserID            int        `db:"userId" json:"userId"`
	EngineerID        int        `db:"engineerId" json:"engineerId"`
	Status            string     `db:"status" json:"status"` // e.g., "active", "inactive", "completed"
	CreatedAt         time.Time  `db:"createdAt" json:"createdAt"`
	EndJobDate        *time.Time `db:"endJobDate" json:"endJobDate"` // Nullable field
	PayDate           *time.Time `db:"payDate" json:"payDate"`       // Nullable field
	Amount            float64    `db:"amount" json:"amount"`
	ImageBeforeRepair string     `db:"imageBeforeRepair" json:"imageBeforeRepair"`
	ImageAfterRepair  string     `db:"imageAfterRepair" json:"imageAfterRepair"`
	Engineer          string     `db:"engineer" json:"engineer"`
}

// Company represents company details
type Company struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Address      string    `db:"address" json:"address"`
	Phone        string    `db:"phone" json:"phone"`
	Email        string    `db:"email" json:"email"`
	FacebookPage string    `db:"facebook_page" json:"facebook_page"`
	TaxCode      string    `db:"tax_code" json:"tax_code"`
	Status       string    `db:"status" json:"status"`       // e.g., "active", "inactive"
	CreatedAt    time.Time `db:"createdAt" json:"createdAt"` // Timestamp for creation
	UpdatedAt    time.Time `db:"updatedAt" json:"updatedAt"` // Timestamp for updates
}

type Section struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	DepartmentId int       `db:"departmentId" json:"departmentId"`
	Department   string    `db:"department" json:"department"`
	Status       string    `db:"status" json:"status"`
	CreatedAt    time.Time `db:"createdAt" json:"createdAt"`
	Users        []User    `json:"users"`
}
