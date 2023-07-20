package handler

import (
	"be-service-saksi-management/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// AccountHandler is REST API handler for Service Account System
type AccountHandler struct {
	AccountUsecase domain.AccountUsecase
}

// GetValidation is handler for get validation
func (ah *AccountHandler) GetValidation(c *fiber.Ctx) error {
	return c.Status(fasthttp.StatusOK).JSON("respon")
}
