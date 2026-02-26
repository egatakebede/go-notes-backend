package main

import (
	"Backend/database"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	database.ConnectDB()

	routes.RegisterNoteRoutes(r)

	r.Run(":9000")

}
