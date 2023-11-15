package routes

import (
	handlers "github.com/jeronimofalavina/config-manager/cmd/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/configs", handlers.ListConfigs)
	r.POST("/configs", handlers.CreateConfigs)
	r.GET("/configs/:name", handlers.GetConfig)
	r.PUT("/configs/:name", handlers.UpdateConfig)
	r.PATCH("/configs/:name", handlers.UpdateConfig)
	r.DELETE("/configs/:name", handlers.DeleteConfig)
	r.GET("/search", handlers.QueryConfigs)
}
