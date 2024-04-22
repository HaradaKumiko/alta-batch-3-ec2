package appController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<img style='display: block; margin: auto;' src='https://cdn.epicstream.com/images/ncavvykf/epicstream/a54b9c16f0f9e2de831b32febc169e734e4ded3d-1920x1080.png?rect=0,36,1920,1008&w=1200&h=630&auto=format'/>")
}
