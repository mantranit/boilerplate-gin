package account

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Account custom select
type Account struct {
	gorm.Model
	Username  string
	Email     string
	FirstName string
	LastName  string
	Birthday  *time.Time
	Gender    string
	Status    string
	LinkedFB  bool
	LinkedTW  bool
	LinkedGG  bool
	CreatedBy string
	UpdatedBy string
}
