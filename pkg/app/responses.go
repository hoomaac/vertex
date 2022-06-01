package app

const (
	UserRegisteredOk    string = "User registered successfully!"
	UserRegisteredNotOk string = "Failed to create the user!"
	UserNotFound        string = "User not found!"
	OtpIsNotValid       string = "OTP is not valid, please send the login request again!"
)

type GeneralResponse struct {
	Status  int
	Message string
}

type RegisterResponse struct {
	Reponse GeneralResponse // Normal response
}

type LoginResponse struct {
	Response GeneralResponse // Normal response
	Code     string          // OTP code
}

type LoginConfirmResponse struct {
	Response GeneralResponse // Normal response
	Token    string          // JWT
}
