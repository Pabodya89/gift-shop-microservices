package main 

import (
	"github.com/gin-gonic/gin"
	"services/inventory-service/database"
	"services/inventory-service/routes"
)

func main() {
	//connect db
	database.ConnectDB()

	//create gin engine
	r := gin.Default()

	//register routes
	routes.RegisterInventoryRoutes(r)

	//run service
	r.Run(":8082")
}