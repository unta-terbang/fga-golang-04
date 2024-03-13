package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "main.go/models"
	"time"
	"net/http"
    "strconv"
)


func CreateOrder(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var requestData struct {
            CustomerName string `json:"customerName"`
            Items        []struct {
                ItemCode    string `json:"itemCode"`
                Description string `json:"description"`
                Quantity    uint   `json:"quantity"`
            } `json:"items"`
        }

        if err := c.ShouldBindJSON(&requestData); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        // Get the current time
        orderedAt := time.Now()

        // Create the order
        order := models.Order{
            Costumer_name: requestData.CustomerName,
            Ordered_at:    orderedAt,
        }
        db.Create(&order)

        // Create items associated with the order
        for _, itemData := range requestData.Items {
            item := models.Item{
                Code:        itemData.ItemCode,
                Description: itemData.Description,
                Quantity:    itemData.Quantity,
                OrderID:     order.ID, // Associate the item with the order
            }
            db.Create(&item)
        }

        c.JSON(200, gin.H{"message": "Order created successfully"})
    }
}


func GetOrderByID(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var order models.Order
        id := c.Param("id")
        if err := db.Preload("Items").First(&order, id).Error; err != nil {
            c.JSON(404, gin.H{"error": "Order not found"})
            return
        }

        // Response structure
        response := struct {
            ID           uint       `json:"id"`
            OrderedAt    time.Time  `json:"orderedAt"`
            CustomerName string     `json:"customerName"`
            Items        []struct {
                ItemCode    string `json:"itemCode"`
                Description string `json:"description"`
                Quantity    uint   `json:"quantity"`
            } `json:"items"`
        }{
            ID:           order.ID,
            OrderedAt:    order.Ordered_at,
            CustomerName: order.Costumer_name,
            Items:        make([]struct {
                ItemCode    string `json:"itemCode"`
                Description string `json:"description"`
                Quantity    uint   `json:"quantity"`
            }, len(order.Items)),
        }

        for i, item := range order.Items {
            response.Items[i] = struct {
                ItemCode    string `json:"itemCode"`
                Description string `json:"description"`
                Quantity    uint   `json:"quantity"`
            }{
                ItemCode:    item.Code,
                Description: item.Description,
                Quantity:    item.Quantity,
            }
        }

        c.JSON(200, response)
    }
}

// func UpdateOnlyOrder(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         var requestData struct {
//             CustomerName string    `json:"customerName"`
//         }

//         // Bind the JSON request body to the requestData struct
//         if err := c.ShouldBindJSON(&requestData); err != nil {
//             c.JSON(400, gin.H{"error": err.Error()})
//             return
//         }

//         // Get the order ID from the URL params
//         id := c.Param("id")

//         // Find the order by ID
//         var order models.Order
//         if err := db.Preload("Items").First(&order, id).Error; err != nil {
//             c.JSON(404, gin.H{"error": "Order not found"})
//             return
//         }
        
//         orderedAt := time.Now()

//         // Update the order fields
//         order.Ordered_at = orderedAt
//         order.Costumer_name = requestData.CustomerName

//         // Save the updated order to the database
//         if err := db.Save(&order).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
//             return
//         }

//         c.JSON(200, order)
//     }
// }

func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var requestData struct {
            OrderedAt    time.Time `json:"orderedAt"`
            CustomerName string    `json:"customerName"`
            Items        []struct {
                ItemCode    string `json:"itemCode"`
                Description string `json:"description"`
                Quantity    uint   `json:"quantity"`
            } `json:"items"`
        }

        // Bind the JSON request body to the requestData struct
        if err := c.ShouldBindJSON(&requestData); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        // Get the order ID from the URL params
        id := c.Param("id")
        // Convert id to uint
        orderID, err := strconv.ParseUint(id, 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
            return
        }

        // Find the order by ID
        var order models.Order
        if err := db.Preload("Items").First(&order, orderID).Error; err != nil {
            c.JSON(404, gin.H{"error": "Order not found"})
            return
        }

        // Delete existing items associated with the order
        if err := db.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
            return
        }
        

        // Update the order fields
        order.Ordered_at = requestData.OrderedAt
        order.Costumer_name = requestData.CustomerName

        // Create new items associated with the order
        for _, itemData := range requestData.Items {
            item := models.Item{
                Code:        itemData.ItemCode,
                Description: itemData.Description,
                Quantity:    itemData.Quantity,
                OrderID:     uint(orderID),
            }
            db.Create(&item)
        }

        // Save the updated order to the database
        db.Save(&order)

        c.JSON(200, order)
    
    }
}

// func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         var requestData struct {
//             OrderedAt    time.Time `json:"orderedAt"`
//             CustomerName string    `json:"customerName"`
//             Items        []struct {
//                 ItemCode    string `json:"itemCode"`
//                 Description string `json:"description"`
//                 Quantity    uint   `json:"quantity"`
//             } `json:"items"`
//         }

//         // Bind the JSON request body to the requestData struct
//         if err := c.ShouldBindJSON(&requestData); err != nil {
//             c.JSON(400, gin.H{"error": err.Error()})
//             return
//         }

//         // Get the order ID from the URL params
//         id := c.Param("id")


//         // Find the order by ID
//         var order models.Order
//         if err := db.Preload("Items").First(&order, id).Error; err != nil {
//             c.JSON(404, gin.H{"error": "Order not found"})
//             return
//         }

//         // Delete existing items associated with the order
//         if err := db.Where("order_id = ?", order.ID).Delete(&models.Item{}).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
//             return
//         }

//         // Update the order fields
//         order.Ordered_at = requestData.OrderedAt
//         order.Costumer_name = requestData.CustomerName

//         // Create new items associated with the order
//         for _, itemData := range requestData.Items {
//             item := models.Item{
//                 Code:        itemData.ItemCode,
//                 Description: itemData.Description,
//                 Quantity:    itemData.Quantity,
//                 OrderID:     order.ID,
//             }
//             db.Create(&item)
//         }

//         // Save the updated order to the database
//         db.Save(&order)

//         c.JSON(200, order)
//     }
// }



// func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         var requestData struct {
//             OrderedAt    time.Time `json:"orderedAt"`
//             CustomerName string    `json:"customerName"`
//             Items        []struct {
//                 ItemCode    string `json:"itemCode"`
//                 Description string `json:"description"`
//                 Quantity    uint   `json:"quantity"`
//             } `json:"items"`
//         }

//         // Bind the JSON request body to the requestData struct
//         if err := c.ShouldBindJSON(&requestData); err != nil {
//             c.JSON(400, gin.H{"error": err.Error()})
//             return
//         }

//         // Get the order ID from the URL params
//         id := c.Param("id")

//         // Find the order by ID
//         var order models.Order
//         if err := db.Preload("Items").First(&order, id).Error; err != nil {
//             c.JSON(404, gin.H{"error": "Order not found"})
//             return
//         }

//         // Delete existing items associated with the order
//         if err := db.Where("order_id = ?", order.ID).Delete(&models.Item{}).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
//             return
//         }

//         // Update the order fields
//         order.Ordered_at = requestData.OrderedAt
//         order.Costumer_name = requestData.CustomerName

//         // Create new items associated with the order
//         for _, itemData := range requestData.Items {
//             item := models.Item{
//                 Code:        itemData.ItemCode,
//                 Description: itemData.Description,
//                 Quantity:    itemData.Quantity,
//                 OrderID:     order.ID,
//             }
//             db.Create(&item)
//         }

//         // Save the updated order to the database
//         db.Save(&order)

//         c.JSON(200, order)
//     }
// }


func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the order ID from the URL params
        id := c.Param("id")

        // Find the order by ID
        var order models.Order
        if err := db.Preload("Items").First(&order, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
            return
        }

        // Delete associated items
        if err := db.Where("order_id = ?", order.ID).Delete(&models.Item{}).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
            return
        }

        // Delete the order
        if err := db.Delete(&order).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Order and associated items deleted successfully"})
    }
}

