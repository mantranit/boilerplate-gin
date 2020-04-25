package account

import (
	"time"
)

// Account custom select
type Account struct {
	// gorm.Model
	ID        uint       `gorm:"primary_key" json:"accountId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`

	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Birthday  *time.Time `json:"birthday"`
	Gender    string     `json:"gender"`
	Status    string     `json:"status"`
	LinkedFB  bool       `json:"linkedFB"`
	LinkedTW  bool       `json:"linkedTW"`
	LinkedGG  bool       `json:"linkedGG"`
	CreatedBy string     `json:"createdBy"`
	UpdatedBy string     `json:"updatedBy"`
}
