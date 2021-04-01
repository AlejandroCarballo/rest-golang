package transport

import (
	"net/http"

	"github.com/AlejandroCarballo/ejemplo/internal/service"
	"github.com/labstack/echo"
)

type HTTPTransport interface {
	Run() error
}

type httpTransport struct {
	is service.ItemService
	e  *echo.Echo
}

func NewHTTP(is service.ItemService) HTTPTransport {
	e := echo.New()
	e.GET("/", runHelloWold)
	return &httpTransport{is, e}
}

func runHelloWold(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (ht *httpTransport) Run() error {
	return ht.e.Start("0.0.0.0:1234")
}
