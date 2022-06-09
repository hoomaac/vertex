package app

type LoginRequest struct {
	Email string `json:"email"`
}

type LoginConfirmRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"` // OTP code
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
