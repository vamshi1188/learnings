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
func (r *OrderRepository) InsertData() error {

	id := 3
	product := ""
	ammount := 1
	_, err := r.DB.Exec(`INSERT INTO orders(id,product,amount)VALUES($1,$2,$3)`, id, product, ammount)
	return err
}
func (r *OrderRepository) DeleteData() error {

	id := 3

	_, err := r.DB.Exec(`DELETE FROM orders WHERE id = $1;`, id)
	return err
}
