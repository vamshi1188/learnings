package main

import (
	"database/sql"
	"log"
	"psql/database"

	// "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5433 user=postgres password=saivamshi88 dbname=mydb sslmode=disable"

	dbConnection, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()
	OrderRepository := &database.OrderRepository{DB: dbConnection}
	err = OrderRepository.CreateTable()
	if err != nil {
		log.Fatal("error in connecting database", err)
	}
	err = OrderRepository.DeleteData()
	if err != nil {
		log.Fatal("error in inserting data in table ", err)
	}

}
