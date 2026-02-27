package routes

import (
	"backend/controllers"
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterNoteRoutes(r *gin.Engine) {
	// Notes routes
	r.GET("/notes", getNotes)
	r.POST("/notes", createNote)
	r.PUT("/notes/:id", updateNote)
	r.DELETE("/notes/:id", deleteNote)

	// User authentication routes
	r.POST("/users/signup", controllers.SignUp)
	r.POST("/users/login", controllers.SignIn)
}

func getNotes(c *gin.Context) {
	var notes []models.Note
	search := c.Query("search")

	if search != "" {
		database.DB.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%").Find(&notes)
	} else {
		database.DB.Find(&notes)
	}

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

func updateNote(c *gin.Context) {
	id := c.Param("id")

	var note models.Note
	if err := database.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var input models.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.Title = input.Title
	note.Content = input.Content
	database.DB.Save(&note)

	c.JSON(http.StatusOK, note)
}

func deleteNote(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Note{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
