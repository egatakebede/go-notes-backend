package controllers

import (
	"net/http"

	"gtihub.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "frontend/signup.html", gin.H{})

}
func SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "frontend/SignIn.html", gin.H{})
}
