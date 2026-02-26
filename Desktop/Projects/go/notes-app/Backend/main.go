package main

import (
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.connectDB()

	routes.RegisterNoteRoutes(r)
	r.Run(":6000")

}
