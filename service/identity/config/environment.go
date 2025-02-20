package config

type environment struct {
	ServerPort string `env:"SERVER_PORT,required"`
	ServerHost string `env:"SERVER_HOST,required"`
}
