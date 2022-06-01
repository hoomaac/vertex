package config

type DataBaseConfig struct {
	Username string
	Password string
	Ip       string
	Port     string
	Name     string
}

type RedisConfig struct {
	Ip       string
	Port     string
	Password string
	Database int
}
