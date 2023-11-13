package routes

import (
	handlers "HFtest-platform-jeronimofalavina-tst/cmd/handler"

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
