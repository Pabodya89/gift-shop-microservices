package route 

import (
	"github.com/gin-gonic/gin"
	"services/inventory-service/internal/handler"
	"services/inventory-service/internal/repository"
	"services/inventory-service/internal/service"
)

func RegisterInventoryRoutes(r *gin.Engine) {
	itemRepo := repository.NewItemRepository()
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	inventoryGroup := r.Group("/inventory")
	{
		inventoryGroup.POST("/items", itemHandler.CreateItem)
		inventoryGroup.GET("/items/:id", itemHandler.GetItemByID)
		inventoryGroup.PUT("/items", itemHandler.UpdateItem)
		inventoryGroup.DELETE("/items/:id", itemHandler.DeleteItem)
	}
}