package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Global template variable - parse once at startup
var htmlTemplate *template.Template

// LoadTemplate renders HTML files
func LoadTemplate(c *gin.Context) {
	// Execute template directly to the response writer
	err := htmlTemplate.Execute(c.Writer, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error while executing template")
		return
	}

	// Set content type for HTML
	c.Header("Content-Type", "text/html; charset=utf-8")
}

func main() {
	// Parse template once at startup
	var err error
	htmlTemplate, err = template.ParseFiles("web.html")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create Gin router
	app := gin.Default()
	fmt.Println("Welcome to my Gin server..")

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Go to http://localhost:8080/count")
	})

	app.GET("/count", LoadTemplate)

	app.GET("/clickme", func(c *gin.Context) {
		c.String(http.StatusOK, "you have clicked")
	})

	fmt.Println("Server starting on http://localhost:8080")
	app.Run(":8080")
}
