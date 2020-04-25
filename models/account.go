package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Account object
type Account struct {
	gorm.Model
	Username  string `gorm:"unique" binding:"required"`
	Email     string `gorm:"unique" binding:"required"`
	Hash      string `binding:"required"`
	FirstName string
	LastName  string
	Birthday  *time.Time
	Gender    string // FEMALE, MALE, OTHER
	Status    string // ACTIVE, DEACTIVATED, PENDING, LOCKED
	LinkedFB  bool   `gorm:"default=false"`
	LinkedTW  bool   `gorm:"default=false"`
	LinkedGG  bool   `gorm:"default=false"`
	CreatedBy string
	UpdatedBy string
	Token     string
}
