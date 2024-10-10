package routes

import (
	"api-mongo-go/internal/controllers"
	"api-mongo-go/internal/core/providers"
)

func SetupRoutes(provider *providers.ApiProvider) {
	mcServerController := controllers.NewMinecraftServerController(provider.DB)

	servers := provider.Router.Group("/servers")
	{
		servers.POST("", mcServerController.CreateServer)
		servers.GET("", mcServerController.GetAllServers)
		servers.GET("/:id", mcServerController.GetServer)
		servers.PATCH("/:id", mcServerController.UpdateServer)
		servers.DELETE("/:id", mcServerController.DeleteServer)
	}

}
