package main 

import (
	"github.com/gin-gonic/gin"
	"services/order-service/database"
	"services/order-service/internal/routes"
)

func main() {
	//connect db
	database.ConnectDB()

	//create gin engine
	r := gin.Default()

	//register routes
	routes.RegisterOrderRoutes(r)

	//run service
	r.Run(":8083")
}