package main

import (
	"providers/handlers"
	"providers/repositories"
	"providers/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	r := repositories.NewDefaultCustomerRepository()
	s := services.NewCustomerService(r)
	h := handlers.NewCustomerHandler(s)

	e.GET("/customers", h.FindAll)

	e.Start(":4040")
}
