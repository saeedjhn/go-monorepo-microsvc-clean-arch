package http

import (
	"fmt"
	"net/http"

	"github.com/saeedjhn/go-monorepo-microsvc-clean-arch/configs/user"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Cfg    *user.Config
	Router *echo.Echo
}

func New(cfg *user.Config) *Server {
	return &Server{
		Cfg:    cfg,
		Router: echo.New(),
	}
}

func (s Server) Run() error {
	// Router Setup
	// http://localhost:80/v1/users
	s.Router.GET("/v1/users/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world!!-----")
	})

	// http://localhost:80/v1/users/about
	s.Router.GET("/v1/users/about", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "About page")
	})

	return s.Router.Start(fmt.Sprintf(":%s", s.Cfg.HTTPServer.Port))
}
