package auth

// FormLogin object
type FormLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// FormRegister object
type FormRegister struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Type     string `json:"type" binding:"required"`
}
