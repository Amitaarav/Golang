package routes

import (
	"github.com/gin-gonic/gin"
	"golang-restaurantmanagementapp/controllers"

)

func FoodRouter(incomingRoutes *gin.Engine) {
    incomingRoutes.GET("/food",controller.GetFoods())
	incomingRoutes.GET("/food/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods",controller.CreateFood())
	incomingRoutes.PATCH("/food/:food_id",controller.UpdateFood())

}