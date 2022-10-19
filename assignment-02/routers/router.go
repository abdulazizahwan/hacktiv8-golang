package routers

import (
	"assignment-02/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetAllOrders)
	router.POST("/orders", controllers.PostOrders)
	router.DELETE("/orders/:orderID", controllers.DeleteOrderById)
	router.PUT("/orders/:orderID", controllers.UpdateOrderByID)

	return router
}
