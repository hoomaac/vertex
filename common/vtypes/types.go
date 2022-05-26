package vtypes

const (
	Ok           int = 200
	Accepted     int = 202
	BadRequest   int = 400
	Unauthorized int = 401
	NotFound     int = 404
)

type AuthResponse struct {
	Status int
	Data   interface{}
}

type DataBaseInfo struct {
	Username string
	Password string
	Ip       string
	Port     string
	Name     string
}
