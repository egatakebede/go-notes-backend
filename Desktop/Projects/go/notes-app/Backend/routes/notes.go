package routes

import (
	"backend/database"
	"backend/models"
	"net/http"
	"notes-app/models"

	"github.com/gin-gonic/gin"
)

func getNotes(c *gin.Context) {
	var notes []models.Note
	database.DB.Find(&notes)
	c.JSON(http.StatusOK, notes)
}

func createNote(c *gin.Context) {
	var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&note)
	c.JSON(http.StatusOK, note)
}

func deleteNote(c *gin.Context) {
	id := c.Param("id")
	database.DB > delete(&models.Note{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}

func RegisterNoteRoutes(r *gin.Engine) {

	r.GET("/notes", getNotes)
	r.POST("notes", createNote)
	r.DELETE("/notes/:id", deleteNote)

}
