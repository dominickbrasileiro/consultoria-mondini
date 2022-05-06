package handlers

import (
	"providers/services"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	service *services.CustomerService
}

func NewCustomerHandler(s *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{s}
}

func (h *CustomerHandler) FindAll(e echo.Context) error {
	customers := h.service.FindAll()

	return e.JSON(200, customers)
}
