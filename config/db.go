package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/models"
	"main.go/routers"
	"github.com/gin-gonic/gin"
)

var (
	host = "localhost"
	dbPort = "5432"
	user = "postgres"
	password = "root"
	dbName = "tugas-kedua"
	db *gorm.DB
	err error
)

func StartDB() {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error konek ke database: ", err)
	}

	db.Debug().AutoMigrate(models.Item{}, models.Order{})

	router := gin.Default()
    routers.SetupRoutes(router, db)

    router.Run(":8080")
}

func GetDB() *gorm.DB {
	return db
}

// func GetOrderByID(c *gin.Context) {
//     var order models.Order
//     id, err := strconv.Atoi(c.Param("id"))
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
//         return
//     }
//     if err := models.db.First(&order, id).Error; err != nil {
//         c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
//         return
//     }
//     c.JSON(http.StatusOK, order)
// }

// func GetOrderByID(c *gin.Context) {
//     var order models.Order
//     id, err := strconv.Atoi(c.Param("id"))
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
//         return
//     }
//     if err := models.db.First(&order, id).Error; err != nil {
//         c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
//         return
//     }
//     c.JSON(http.StatusOK, order)
// }