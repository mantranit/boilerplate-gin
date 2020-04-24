package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Account object
type Account struct {
	gorm.Model
	Username  string     `gorm:"unique" json:"username" binding:"required"`
	Email     string     `gorm:"unique" json:"email" binding:"required"`
	Hash      string     `binding:"required"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Birthday  *time.Time `json:"birthday"`
	Gender    string     `json:"gender"`
	LinkedFB  bool       `gorm:"default=false"`
	LinkedGG  bool       `gorm:"default=false"`
	CreatedBy string     `json:"createdBy"`
	UpdatedBy string     `json:"updatedBy"`
}
