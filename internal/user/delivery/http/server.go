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
	s.Router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})

	return s.Router.Start(fmt.Sprintf(":%s", s.Cfg.HTTPServer.Port))
}
