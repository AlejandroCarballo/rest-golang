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
	e.GET("/items/:id", getItems)
	e.POST("/users", saveItem)
	return &httpTransport{is, e}
}

func getItems(c echo.Context) error {
	item := c.Param("id")
	return c.String(http.StatusOK, item)

}

func saveItem(c echo.Context) error {
	item := c.FormValue("item")
	description := c.FormValue("description")
	return c.String(http.StatusOK, "item:"+item+", description:"+description)
}

func (ht *httpTransport) Run() error {
	return ht.e.Start("0.0.0.0:1234")
}
