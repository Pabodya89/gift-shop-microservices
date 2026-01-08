package handler

import (
	"inventory-service/services/inventory-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	items :=  service.GetAllItems()
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	idStr := Param("id")
	
}