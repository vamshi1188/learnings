package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message ": "success",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("error in  running server on port 8080", err)
	}
}
