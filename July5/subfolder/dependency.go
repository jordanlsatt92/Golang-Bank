package subfolder

// Jordan Satterfield
import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// add array of transactions
type BankAccount struct{
	Name string
	AccountNum int
	Balance float64
	Transactions [] *Transaction
}

type Transaction struct{
	Time string
	Date string
	Amount float64
}

// deposit; negative should show error
func (a *BankAccount) Deposit(amount float64) float64{
	if amount < 0.0 {
		fmt.Println("You cannot deposit a negative amount of money.")
		return 0
	}else{
		currentTime:=time.Now()
		var date string = (currentTime.Month().String()) + " " + fmt.Sprint(currentTime.Day()) + " " + fmt.Sprint(currentTime.Year())
		time:=fmt.Sprint(currentTime.Hour()) + ":" + fmt.Sprint(currentTime.Minute())
		a.Balance += amount
		a.Transactions = append(a.Transactions, &Transaction{time, date, amount})
		return a.Balance
	}
}
// withdrawel; must maintain minimun balance
func (a *BankAccount) Withdrawel(amount float64) float64{
	if a.Balance - amount < 0{
		fmt.Println(a.Name,"\nYou cannot withdraw an amount greater than your balance")
		fmt.Println("Your account balance:", a.Balance, "\nYour attempted withdrawel:", amount)
		return a.Balance
	}else{
		currentTime:=time.Now()
		var date string = (currentTime.Month().String()) + " " + fmt.Sprint(currentTime.Day()) + " " + fmt.Sprint(currentTime.Year())
		time:= fmt.Sprint(currentTime.Hour()) + ":" + fmt.Sprint(currentTime.Minute())
		a.Balance -= amount
		a.Transactions = append(a.Transactions, &Transaction{time, date, -amount})
		return a.Balance
	}
}

func GenerateAccount() func() int{
	var number int = 1
	var temp int
	return func() int {
		temp = number
		number += 1
		return temp
	}
}

func SortByName(array []*BankAccount){
	sort.Slice(array, func(i, j int) bool {
		str1:= strings.Split(array[i].Name, " ")
		name1:= str1[0]
		str2:= strings.Split(array[j].Name, " ")
		name2:= str2[0]
		return name1 < name2
	})
}

func SortByBalance(array []*BankAccount){
	sort.Slice(array, func(i, j int) bool{
		return array[i].Balance < array[j].Balance
	})
}