package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type VehicleDatabase struct {
	DB *sql.DB
}
type VehiclesTable struct {
	id            int
	vehicleNumber string
	vehicleType   string
	ownerName     string
}

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

	vehicleDB := &VehicleDatabase{DB: dbConnection}

	vehicles, err := vehicleDB.GetVehicles()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range vehicles {

		fmt.Println(v)
	}

}

func (r *VehicleDatabase) GetVehicles() ([]VehiclesTable, error) {

	rows, err := r.DB.Query(`SELECT id, vehicle_number, vehicle_type, owner_name
		FROM vehicles`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehiclesData1 []VehiclesTable

	for rows.Next() {

		var v VehiclesTable

		err := rows.Scan(&v.id, &v.vehicleNumber, &v.vehicleType, &v.ownerName)
		if err != nil {

			return nil, err
		}

		vehiclesData1 = append(vehiclesData1, v)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return vehiclesData1, nil

}
