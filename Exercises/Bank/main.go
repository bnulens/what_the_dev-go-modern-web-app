package main

import (
	"fmt"

	"github.com/bnulens/bank-example/fileoperations"
	"github.com/bnulens/bank-example/messaging"
)

const account string = "XYZ Balance.txt"

func checkBalance(currentBalance *float64) {
	fmt.Printf("Your current balance is: %.2f \n", *currentBalance)
	fmt.Print("-------------------------------\n")
}

func withdrawMoney(currentBalance *float64, amountToWithDraw float64) float64 {
	for *currentBalance < amountToWithDraw {
		fmt.Printf("Cannot withdraw more than %v \n", *currentBalance)
		fmt.Print("Fill in a lower amount: \n")
		fmt.Scan(&amountToWithDraw)
	}
	for amountToWithDraw <= 0 {
		fmt.Printf("Cannot withdraw a negative amount")
		fmt.Print("Fill in another amount: \n")
		fmt.Scan(&amountToWithDraw)
	}
	fmt.Print("Withdrawal successful, you will be returned shortly.\n")
	fmt.Print("----------------------------------------------------\n")

	*currentBalance -= amountToWithDraw
	checkBalance(currentBalance)
	return *currentBalance
}

func depositMoney(balance *float64, deposit float64) {
	*balance += deposit
	checkBalance(balance)
}

func main() {
	var choice int

	var withdrawalAmount float64
	var depositAmount float64

	fmt.Println("Welcome to XYZ-bank !")
	fmt.Println("---------------------")

	accountBalance := fileoperations.GetBalanceFromFile(account)

	for choice < 4 {
		messaging.PrintChoices()
		fmt.Print("Enter a number: ")
		fmt.Scan(&choice)

		toCheckBalance := choice == 1
		toDepositMoney := choice == 2
		toWithdrawMoney := choice == 3

		fmt.Printf("You entered: %v \n", choice)
		if toCheckBalance {
			checkBalance(&accountBalance)
		} else if toDepositMoney {
			fmt.Print("Enter an amount to deposit into your account: ")
			fmt.Scan(&depositAmount)
			depositMoney(&accountBalance, depositAmount)
		} else if toWithdrawMoney {
			fmt.Print("Enter an amount to withdraw from your account: ")
			fmt.Scan(&withdrawalAmount)
			withdrawMoney(&accountBalance, withdrawalAmount)
		} else {
			fmt.Print("Invalid choice \n")
			fmt.Print("Pick another option \n")
		}
	}
	fileoperations.WriteBalanceToFile(&accountBalance)
	fmt.Println("Session ended")
}
