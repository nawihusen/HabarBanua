package http

import (
	"be-service-saksi-management/domain"
	"be-service-saksi-management/saksi/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// RouterAPI is main router for this Service Saksi REST API
func RouterAPI(app *fiber.App, account domain.AccountUsecase) {
	handlerAccount := &handler.AccountHandler{AccountUsecase: account}

	basePath := viper.GetString("server.base_path")
	saksi := app.Group(basePath)

	// Account management

	saksi.Get("/saksi/validation", handlerAccount.GetValidation)

}
