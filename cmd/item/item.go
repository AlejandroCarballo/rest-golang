package main

import (
	"github.com/juanpablopizarro/ejemplo/internal/service"
	"github.com/juanpablopizarro/ejemplo/internal/transport"
)

func main() {
	is := service.NewItemService()
	http := transport.NewHTTP(is)
	http.Run()
}
