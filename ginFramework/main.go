package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message ": "success",
		})
	})

	router.GET("/me/:id", func(c *gin.Context) {

		var id = c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})
	})

	router.POST("/me", func(c *gin.Context) {

		type meREQUEST struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var meRequest meREQUEST

		c.BindJSON(&meRequest)

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})

	})

	router.Run(":8080")

}
