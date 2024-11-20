package config

type (
	Config struct {
		Server *Server
		DB     *Db
	}
	Server struct {
		Port int
	}
	Db struct {
		Host      string
		Port      int
		User      string
		Password  string
		DB_Name   string
		SSL_Name  string
		Time_Zone string
	}
)

type Enviroment struct {
	DB_HOST     string `mapstructure:"DB_HOST" validate:"required"`
	DB_PORT     int    `mapstructure:"DB_PORT" validate:"required"`
	DB_DATABASE string `mapstructure:"DB_DATABASE" validate:"required"`
	DB_USERNAME string `mapstructure:"DB_USERNAME" validate:"required"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD" validate:"required"`
	SERVER_PORT int    `mapstructure:"SERVER_PORT" validate:"required"`
}
