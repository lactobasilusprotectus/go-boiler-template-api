package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lactobasilusprotectus/go-template/config"
	"log"
	"net/http"
	"time"
)

type Server struct {
	port   string
	server http.Server
	echo   *echo.Echo
}

type RouteHandler interface {
	Register(*echo.Echo)
}

//NewServer => merupakan construct dari struct Server
func NewServer(cfg config.HttpConfig) *Server {
	e := echo.New()

	e.Use(middleware.CORS())

	return &Server{
		port: cfg.Port,
		server: http.Server{
			Addr:         fmt.Sprintf("localhost:%s", cfg.Port),
			WriteTimeout: time.Millisecond * time.Duration(cfg.Timeout),
			ReadTimeout:  time.Millisecond * time.Duration(cfg.Timeout),
			IdleTimeout:  time.Second * 60,
		},
		echo: e,
	}
}

//RegisterHandler => fungsi untuk mendaftar Api Handler untuk bisa di gunakan dengan echo framework
func (s *Server) RegisterHandler(route RouteHandler) {
	route.Register(s.echo)
}

//Run => fungsi untuk menjalankan echo di port tertentu
func (s *Server) Run(env string) {
	go func() {
		log.Println("starting", env, "server on port", s.port)

		if err := s.echo.StartServer(&s.server); err != nil {
			log.Fatal(err)
		}
	}()
}

//Stop => fungsi untuk memberhentikan echo
func (s *Server) Stop() {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)

	defer cancelFunc()

	_ = s.echo.Shutdown(timeout)
}
