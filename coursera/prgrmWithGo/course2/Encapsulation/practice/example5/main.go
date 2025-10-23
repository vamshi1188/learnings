package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BankAccount struct {
	name          string
	accountnumber int
	balance       float64
}

func (ba *BankAccount) AccountInit(name string, accountNumber int, balance float64) {
	if name == "" || accountNumber == 0 || balance <= 0 {
		fmt.Println("invalid input!")
		return
	}
	ba.name = name
	ba.accountnumber = accountNumber
	ba.balance = balance
}

func (ba *BankAccount) WithdrawMoney(amount float64) {

	if amount <= 0 {
		fmt.Println("inavlid amount!")
		return
	}
	if amount > ba.balance {
		fmt.Println("Insuffient funds in your account!")
		return
	}
	ba.balance -= amount

	fmt.Printf("%v withdrawn from account sucessfully!\n", amount)
}

func (ba *BankAccount) Deposits(amount float64) {

	if amount <= 0 {
		fmt.Println("invalid Deposite amount!")
		return
	}
	ba.balance += amount

	fmt.Printf("%v deposited sucessfully\n", amount)
}

func (ba BankAccount) GetDetails() {

	fmt.Printf("Name := %v\n", ba.name)
	fmt.Printf("AccountNumber := %v\n", ba.accountnumber)
	fmt.Printf("Balance := %v\n", ba.balance)
}
func (ba BankAccount) CheckBalance() {
	fmt.Printf("available balance is %v\n", ba.balance)
}
func WelcomeGreet() {

	fmt.Println("Welcome to Bank of india")
	fmt.Println("choose the options below!")
	fmt.Println("1.get account details")
	fmt.Println("2.Deposit")
	fmt.Println("3.Withdraw")
	fmt.Println("4.Checkbalance")
	fmt.Println("5.Exit")

}

func main() {
	var BankAccount BankAccount
	BankAccount.AccountInit("saivamshi", 5740, 2000)

	sc := bufio.NewScanner(os.Stdin)
	for {
		WelcomeGreet()
		sc.Scan()
		input := sc.Text()

		switch input {
		case "1":
			BankAccount.GetDetails()

		case "2":
			fmt.Println("enter deposit amount")
			sc.Scan()
			DepositeAmount := sc.Text()

			Dpamount, err := strconv.ParseFloat(DepositeAmount, 64)
			if err != nil {
				fmt.Println("Deposit failed!")
				break
			}
			BankAccount.Deposits(Dpamount)

		case "3":
			fmt.Println("enter amount to withdraw")
			sc.Scan()
			withdrawAmount := sc.Text()

			WithdrawAmount, err := strconv.ParseFloat(withdrawAmount, 64)
			if err != nil {
				fmt.Println("Withdraw failed!")
				return
			}
			BankAccount.WithdrawMoney(WithdrawAmount)

		case "4":
			BankAccount.CheckBalance()
		case "5":
			fmt.Println("Thank you for banking with us!")
			os.Exit(0)

		default:
			fmt.Println("Invalide option")
		}

	}
}
