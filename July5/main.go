package main

//Jordan Satterfield
import (
	"fmt"
	"mymodule/subfolder"
)
func main(){

	accounts:=[]*subfolder.BankAccount{}
	f:= subfolder.GenerateAccount()
	var input string
	var firstName string
	var lastName string
	var accountNum int
	var balance float64
	for input != "done"{
		fmt.Print("First name of account holder (string):")
		fmt.Scanln(&firstName)
		fmt.Print("Last name of account holder (string):")
		fmt.Scanln(&lastName)
		accountNum = f()
		fmt.Print("Initial deposit (float):")
		fmt.Scanln(&balance)
		fmt.Print("Type \"done\" if you do not want create more accounts, hit enter if you do:")
		fmt.Scanln(&input)
		a:=subfolder.BankAccount{Name:firstName + " " + lastName, AccountNum:accountNum, Balance:balance}
		accounts = append(accounts, &a)
	}

	fmt.Println("All Accounts:")
	for i:= 0; i < len(accounts);i++{
		fmt.Println(*accounts[i])
	}

	subfolder.SortByName(accounts)
	fmt.Println("\nSorted by first name:")
	for i:= 0; i < len(accounts);i++{
		fmt.Println(*accounts[i])
	}

	subfolder.SortByBalance(accounts)
	fmt.Println("\nSorted by account balance:")
	for i:= 0; i < len(accounts);i++{
		fmt.Println(*accounts[i])
	}

	// Transactions
	for i:= 0; i < len(accounts);i++{
		accounts[i].Deposit(50.66)
		accounts[i].Withdrawel(30.32)
		accounts[i].Deposit(100.32)
		accounts[i].Withdrawel(75.88)
		accounts[i].Deposit(10.21)
		accounts[i].Withdrawel(300)
	}

	fmt.Println("\nTransactions for each user:")
	for i:=0; i < len(accounts);i++{
		fmt.Printf("%s's transactions\n", accounts[i].Name)
		for j:=0; j < len(accounts[i].Transactions); j++{
			fmt.Printf("Time: %s, Date: %s, Transaction: %.2f\n",accounts[i].Transactions[j].Time, accounts[i].Transactions[j].Date, accounts[i].Transactions[j].Amount)
		}
	}

	fmt.Println("\nEnding balances:")
	for i:= 0; i < len(accounts);i++{
		fmt.Printf("%s's ending balance:%.2f\n",accounts[i].Name, accounts[i].Balance)
	}
}