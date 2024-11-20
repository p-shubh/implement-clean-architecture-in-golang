package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"main.go/config"
	"main.go/connections/database"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf config.Config
}

func EchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: *conf,
	}

}

// Start implements Server.
func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.GET("/healthCheck", s.healthCheck)
	s.routes()
	serverUrl := fmt.Sprintf(":%d", s.conf.Env.SERVER_PORT)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.JSON(200, "server is up")
}

func (s *echoServer) routes() {

}
