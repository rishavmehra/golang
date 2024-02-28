package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:tables_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:tables_id", controller.UpdateTable())
}
