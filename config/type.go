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
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     int    `mapstructure:"DB_PORT"`
	DB_DATABASE string `mapstructure:"DB_DATABASE"`
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
}
