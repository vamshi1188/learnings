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
func (r *OrderRepository) InsertData(order Order) error {

	_, err := r.DB.Exec(`INSERT INTO orders(product,amount)VALUES($1,$2)`, order.Product, order.Amount)
	return err
}
func (r *OrderRepository) DeleteData(id int) error {

	_, err := r.DB.Exec(`DELETE FROM orders WHERE id = $1;`, id)
	return err
}
func (r *OrderRepository) GetAll() ([]Order, error) {

	rows, err := r.DB.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.Id, &order.Product, &order.Amount)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) GetById(id int) (Order, error) {
	var order Order

	err := r.DB.QueryRow("SELECT * FROM orders WHERE id = $1", id).Scan(&order.Id, &order.Product, &order.Amount)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func (r *OrderRepository) UpdateRow(order Order) error {

	_, err := r.DB.Exec("UPDATE orders SET product = $1,amount = $2 WHERE id = $3", order.Product, order.Amount, order.Id)
	return err

}
