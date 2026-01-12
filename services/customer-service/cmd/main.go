package main

import (
	"github.com/gin-gonic/gin"
	"services/customer-service/database"
	"services/customer-service/internal/routes"
)

func main() {
	//connect db
	database.ConnectDB()

	//create gin engine
	r := gin.Default()

	//register routes
	routes.RegisterCustomerRoutes(r)

	//run service
	r.Run(":8081")
}
