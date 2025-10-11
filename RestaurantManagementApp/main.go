package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"restaurantmanagementapp/database"
	"restaurantmanagementapp/router"
	"restaurantmanagementapp/middlewares"
	"middlewares/Authentication"
	"go.mongodb.org/mongo-driver/mongo"

)

var foodcollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getrnv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middlewares,Authentication());

	routes.FoodRouter(router)
	routes.OrderRouter(router)
	routes.UserRouter(router)
	routes.TableRouter(router)
	routes.OrderRouter(router)
	routes.InvoiceRouter(router)
	routes.OrderItemRouter(router)
	routes.MenuRouter(router)

	router.Run(":" + port)
}