package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RootHandler struct {
	Env       string `json:"env"`
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
}

func NewRootHandler(env string) *RootHandler {
	return &RootHandler{Env: env, Version: "n/a"}
}

func (h *RootHandler) Register(e *echo.Echo) {
	e.GET("/", h.Root)
}

func (h *RootHandler) Root(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, h)
}
