package database

import "database/sql"

type OrderRepository struct {
	DB *sql.DB
}
type Order struct {
	Id      int
	Product string
	Amount  int
}

func (r *OrderRepository) CreateTable() error {

	_, err := r.DB.Exec(`CREATE TABLE IF NOT EXISTS  orders(
	id SERIAL PRIMARY KEY ,
	product TEXT,
	amount INTEGER
	)`)
	return err
}
