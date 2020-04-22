package models

// UserModel class
type UserModel struct{}

// User object
type User struct {
	Email     string `json:"email" binding:"required"`
	Hash      string `json:"hash" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
}
