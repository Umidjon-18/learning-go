package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	AccountOwner  string
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	acc.Balance += amount
	fmt.Printf("Deposit added %.2f to %s. New balance $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be positive")
	}
	if amount > acc.Balance {
		return fmt.Errorf("Insufficient funds in %s. Balance $%.2f. Tried to withdraw $%.2f",
			acc.AccountNumber,
			acc.Balance,
			amount,
		)
	}
	acc.Balance -= amount
	fmt.Printf("Withdrew $%.2f from %s. New balance $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 { return acc.Balance }

func (acc *Account) String() string {
	return fmt.Sprintf("Account %s. Owner %s. Balance %.2f\n",
		acc.AccountNumber,
		acc.AccountOwner,
		acc.Balance,
	)
}

type SavingAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest %.2f to saving account %s.\n", interest, sa.AccountNumber)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Printf("AddIntest: Error depositing $%.2f to saving account. %v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be positive")
	}

	if oa.Balance+oa.OverdraftLimit < amount {
		return fmt.Errorf("Withdrawal of $%.2f is exceeds overdraft limit for %s. Available including overdraft $%.2f",
			amount,
			oa.AccountNumber,
			oa.Balance+oa.OverdraftLimit,
		)
	}
	oa.Balance -= amount
	fmt.Printf("Withdrew $%.2f from overdraft account %s. New balance $%.2f\n",
		amount,
		oa.AccountNumber,
		oa.Balance,
	)
	return nil
}

func main() {
	fmt.Println("------- Bank Account System -------")

	savAcc := SavingAccount{
		Account: Account{
			AccountNumber: "SAV001",
			AccountOwner:  "John Doe",
			Balance:       1000,
		},
		InterestRate: 0.04,
	}

	fmt.Println("------- Saving Account Operations -------")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(400)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	savAcc.AddInterest()
	err = savAcc.Withdraw(400)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	ovrAcc := OverdraftAccount{
		Account: Account{
			AccountNumber: "OAV001",
			AccountOwner:  "Jany Carl",
			Balance:       200,
		},
		OverdraftLimit: 200,
	}

	fmt.Println("------- Overdraft Account Operations -------")
	fmt.Println(ovrAcc.Account.String())

	err = ovrAcc.Deposit(100)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	err = ovrAcc.Withdraw(400)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	err = ovrAcc.Withdraw(400)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Printf("Final Overdraft account details: %s\n", ovrAcc.Account.String())

}
