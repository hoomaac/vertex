package app

type LoginRequest struct {
	Email string
}

type LoginConfirmRequest struct {
	Email string
	Code  string // OTP code
}

type RegisterRequest struct {
	Username string
	Email    string
}
