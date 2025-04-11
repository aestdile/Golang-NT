





package main

import (
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	ID          int
	Amount      float64
	ForBalance  float64
	Timestamp   time.Time
	Description string
}

type Account struct {
	ID            int
	Name          string
	Balance       float64
	AccountNumber string
	Transactions  []Transaction
}

func (acc *Account) Deposit(amount float64, description string) error {
	if amount <= 0 {
		return errors.New("Value must be greater than 0")
	}
	acc.Balance += amount
	acc.Transactions = append(acc.Transactions, Transaction{
		ID:          len(acc.Transactions) + 1,
		Amount:      amount,
		ForBalance:  acc.Balance,
		Timestamp:   time.Now(),
		Description: description,
	})
	return nil
}

func (acc *Account) Withdraw(amount float64, description string) error {
	if amount <= 0 {
		return errors.New("Value must be greater than 0")
	}
	if amount > acc.Balance {
		return errors.New("Finances not enough, increase your finances, please!!!")
	}
	acc.Balance -= amount
	acc.Transactions = append(acc.Transactions, Transaction{
		ID:          len(acc.Transactions) + 1,
		Amount:      -amount,
		ForBalance:  acc.Balance,
		Timestamp:   time.Now(),
		Description: description,
	})
	return nil
}

func (acc *Account) PrintStatement() {
	fmt.Printf("Statement for Account %s (%s):\n", acc.AccountNumber, acc.Name)
	fmt.Printf("Balance: $%.2f\n", acc.Balance)
	fmt.Println("Transactions:")
	for _, trans := range acc.Transactions {
		fmt.Printf("ID: %d, Amount: $%.2f, Description: %s, Timestamp: %s\n", trans.ID, trans.Amount, trans.Description, trans.Timestamp.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	var name, accountNumber string
	var initialBalance float64

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your account number: ")
	fmt.Scanln(&accountNumber)

	fmt.Println("Enter your initial balance: ")
	fmt.Scanln(&initialBalance)

	acc := Account{
		ID:            1,
		Name:          name,
		Balance:       initialBalance,
		AccountNumber: accountNumber,
		Transactions:  []Transaction{},
	}

	if err := acc.Deposit(500, "First deposit: "); err != nil {
		fmt.Println("Deposit error:", err)
	}

	if err := acc.Withdraw(200, "Food shopping: "); err != nil {
		fmt.Println("Withdrawal error:", err)
	}

	acc.PrintStatement()
}












































































