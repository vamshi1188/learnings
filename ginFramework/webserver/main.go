package main

import (
	"database/sql"
	"log"
	"webserver/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	connstr := "host=localhost port=5433 user=vamshi password=saivamshi88 dbname=practicedata "

	dbConnection, err := sql.Open("postgres", connstr)

	if err != nil {

		log.Fatal(err)
	}
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal("database connection failed ", err)
	}
	defer dbConnection.Close()

	vehicleDB := &database.VehicleDatabase{DB: dbConnection}

	router := gin.Default()

	router.GET("/vehicles", func(c *gin.Context) {

		vehicles, err := vehicleDB.FetchVehicles()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(200, vehicles)
	})

	router.Run(":8080")

}
