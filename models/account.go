package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Status of Account
type Status string

// enum Status
const (
	StatusPending     Status = "PENDING"
	StatusActive             = "ACTIVE"
	StatusLocked             = "LOCKED"
	StatusDeactivated        = "DEACTIVATED"
)

// Gender of Account
type Gender string

// enum Status
const (
	GenderOther  Gender = "OTHER"
	GenderFemale        = "FEMALE"
	GenderMale          = "MALE"
)

// Account object
type Account struct {
	gorm.Model
	Email     string `gorm:"unique" binding:"required"`
	Hash      string `binding:"required"`
	FirstName string
	LastName  string
	Birthday  *time.Time
	Gender    Gender
	Status    Status
	LinkedFB  bool `gorm:"default=false"`
	LinkedTW  bool `gorm:"default=false"`
	LinkedGG  bool `gorm:"default=false"`
	CreatedBy string
	UpdatedBy string
	Token     string
}
