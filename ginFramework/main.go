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
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}

		var meRequest meREQUEST

		err := c.BindJSON(&meRequest)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})

	})

	router.GET("/name", func(c *gin.Context) {

	})

	router.Run(":8080")

}
