package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"psql/database"
	"strconv"
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
	for {
		PrintText1()
		// scanner
		sc := bufio.NewScanner(os.Stdin)
		choice := ""
		sc.Scan()
		choice = sc.Text()
		choiceint, err := strconv.Atoi(choice)

		if err != nil {
			log.Fatal("error in string conversion ", err)
		}

		switch choiceint {
		case 1:

			fmt.Println("enter product: ")
			sc.Scan()
			prdct := sc.Text()
			fmt.Printf("enter amount of %v\n", prdct)
			sc.Scan()
			amount := sc.Text()
			amountint, err := strconv.Atoi(amount)
			if err != nil {
				log.Fatal(err)
			}

			err = OrderRepository.InsertData(database.Order{Product: prdct, Amount: amountint})
			if err != nil {
				log.Fatal("error in inserting data in table ", err)
			}

		case 2:
			fmt.Println("enter id:")
			sc.Scan()
			id := sc.Text()
			idint, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("enter product: ")
			sc.Scan()
			prdct := sc.Text()
			fmt.Printf("enter amount of %v\n", prdct)
			sc.Scan()
			amount := sc.Text()
			amountint, err := strconv.Atoi(amount)
			if err != nil {
				log.Fatal(err)
			}
			err = OrderRepository.UpdateRow(database.Order{Id: idint, Product: prdct, Amount: amountint})

		case 3:

			fmt.Println("enter id You want to delete: ")
			sc.Scan()
			id := sc.Text()
			idint, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			err = OrderRepository.DeleteData(idint)

		case 4:
			fmt.Println("enter id: ")
			sc.Scan()
			id := sc.Text()
			idint, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			order, err := OrderRepository.GetById(idint)
			if err != nil {
				log.Fatal("error in getting order by id ", err)
			}
			fmt.Println(order)
		case 5:
			rows, err := OrderRepository.GetAll()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(rows)

		}
	}

}

func PrintText1() {
	fmt.Println()
	fmt.Println("Welcome to myecomerce!")
	fmt.Println("choose one option below : ")
	fmt.Println("1 - for adding product in database")
	fmt.Println("2 - for update product in database")
	fmt.Println("3 - delete product in database ")
	fmt.Println("4 - get data by id")
	fmt.Println("5 - get all data")
}
