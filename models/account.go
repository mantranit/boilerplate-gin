package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Account object
type Account struct {
	gorm.Model
	Username  string     `gorm:"unique" json:"Username" binding:"required"`
	Email     string     `gorm:"unique" json:"Email" binding:"required"`
	Hash      string     `binding:"required"`
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
	Birthday  *time.Time `json:"Birthday"`
	Gender    string     `json:"Gender"` // FEMALE, MALE, OTHER
	Status    string     `json:"Status"` // ACTIVE, DEACTIVATED, PENDING, LOCKED
	LinkedFB  bool       `gorm:"default=false"`
	LinkedTW  bool       `gorm:"default=false"`
	LinkedGG  bool       `gorm:"default=false"`
	CreatedBy string     `json:"CreatedBy"`
	UpdatedBy string     `json:"UpdatedBy"`
}
