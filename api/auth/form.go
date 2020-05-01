package auth

// FormLogin object
type FormLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// FormRegister object
type FormRegister struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

// FormForgotPassword object
type FormForgotPassword struct {
	Email string `json:"email" validate:"required,email"`
}

// FormResetPassword object
type FormResetPassword struct {
	Password        string `form:"password" json:"password" validate:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" validate:"required,eqfield=Password"`
}
