package routes

import (
	"api-wallet/internal/api/wallets"
	"api-wallet/internal/api/wallets/handlers"

	"github.com/pangpanglabs/echoswagger/v2"
)

func RegisterProductRoutes(rootAPI echoswagger.ApiRoot, service wallets.Service) {
	// cfg := configs.GetConfig()
	controller := handlers.NewWalletHandler(service)

	group := rootAPI.Group("wallet", "v1/wallet").SetDescription("Order API")

	group.POST("/verify", controller.PostVerify).
		SetSummary("create verify transaction").
		SetDescription("create verify transaction")

	group.POST("/confirm", controller.PostConfirm).
		SetSummary("confirm transaction").
		SetDescription("confirm transaction")

}
