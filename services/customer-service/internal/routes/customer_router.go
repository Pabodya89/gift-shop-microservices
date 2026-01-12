package routes 

import (
	"github.com/gin-gonic/gin"
	"services/customer-service/internal/handlers"
)

func RegisterCustomerRoutes(r *gin.Engine, customerHandler *handlers.CustomerHandler) {
	customerRoutes := r.Group("/customers")
	{
		customerRoutes.POST("/", customerHandler.CreateCustomer)
		customerRoutes.GET("/:id", customerHandler.GetCustomerByID)
		customerRoutes.PUT("/", customerHandler.UpdateCustomer)
		customerRoutes.DELETE("/:id", customerHandler.DeleteCustomer)
	}
}



