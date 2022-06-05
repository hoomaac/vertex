package app

const (
	UserRegisteredOk    string = "User registered successfully!"
	UserRegisteredNotOk string = "Failed to register the user!"
	UserNotFound        string = "User not found!"
	UserConfirmFailed   string = "User has not been confirmed, please try again!"
	UserConfirmSuccess  string = "User logged-in successfully"
	UserAlreadyConfirm  string = "User is already confirmed!"
	OtpIsNotValid       string = "OTP is not valid, please send the login request again!"
)

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type RegisterResponse struct {
	Reponse GeneralResponse `json:"respones"` // Normal response
}

type LoginResponse struct {
	Response GeneralResponse `json:"response"`       // Normal response
	Code     string          `json:"code,omitempty"` // OTP code
}

type LoginConfirmResponse struct {
	Response GeneralResponse `json:"response"`        // Normal response
	Token    string          `json:"token,omitempty"` // JWT
}
