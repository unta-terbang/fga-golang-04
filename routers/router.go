package routers

import (
    "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/controllers"

)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    router.POST("/orders", controllers.CreateOrder(db))
	router.GET("/orders/:id", controllers.GetOrderByID(db))
	router.PUT("/orders/:id", controllers.UpdateOrder(db))
	// router.PUT("/orders/:id", controllers.UpdateOnlyOrder(db))
	router.DELETE("/orders/:id", controllers.DeleteOrder(db))
}


// func StartServer() *gin.Engine{

// 	router := gin.Default()

// 	router.POST("/items", controllers.CreateItem)

// 	return router

// }


