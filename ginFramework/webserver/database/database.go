package database

import "database/sql"

type VehicleDatabase struct {
	DB *sql.DB
}
type VehiclesTable struct {
	Id            int
	VehicleNumber string
	VehicleType   string
	OwnerName     string
}

func (r *VehicleDatabase) FetchVehicles() ([]VehiclesTable, error) {

	rows, err := r.DB.Query(`SELECT id, vehicle_number, vehicle_type, owner_name
		FROM vehicles`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehiclesData1 []VehiclesTable

	for rows.Next() {

		var v VehiclesTable

		err := rows.Scan(&v.Id, &v.VehicleNumber, &v.VehicleType, &v.OwnerName)
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
