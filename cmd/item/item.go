package main

import (
	"github.com/AlejandroCarballo/ejemplo/internal/service"
	"github.com/AlejandroCarballo/ejemplo/internal/transport"
)

func main() {
	is := service.NewItemService()
	http := transport.NewHTTP(is)
	http.Run()
}
