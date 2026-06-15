package main

import "fmt"

type BankAccount struct {
	Owner         string
	Balance       float64
	AccountNumber string
}

func Deposit(acc *BankAccount, amount float64) error {

	fmt.Println(*acc)

	if amount <= 0 {
		return fmt.Errorf("deposit amount must be greatr than zero")
	} else {
		acc.Balance += amount
		return nil
	}
}

func Withdraw(acc *BankAccount, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be greater than zero")
	} else if acc.Balance < amount {
		return fmt.Errorf("insufficient balance")
	} else {
		acc.Balance -= amount
		return nil
	}
}

func main() {
	fmt.Println("Problem 1")

	var val int
	var amount float64
	acc := &BankAccount{
		Owner:         "Gagandeep Singh",
		Balance:       123520,
		AccountNumber: "23626574854",
	}

	fmt.Println("Enter 1 to Deposit\nEnter 2 to Withdraw")
	fmt.Scan(&val)

	switch val {
	case 1:
		fmt.Println("Enter amount to deposit")
		fmt.Scan(&amount)
		check := Deposit(acc, amount)
		if check != nil {
			fmt.Println("Deposit failed:", check)
		} else {
			fmt.Println("Deposit successful:", acc.Balance)
		}

	case 2:
		fmt.Println("Enter amount to Withdraw")
		fmt.Scan(&amount)
		check := Withdraw(acc, amount)
		if check != nil {
			fmt.Println("Deposit failed:", check)
		} else {
			fmt.Println("Deposit successful:", acc.Balance)
		}
	default:
		return
	}
}
