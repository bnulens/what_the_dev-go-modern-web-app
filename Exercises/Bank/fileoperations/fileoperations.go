package fileoperations

import (
	"fmt"
	"os"
	"strconv"
)

func GetBalanceFromFile(pathString string) float64 {
	data, _ := os.ReadFile(pathString)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func WriteBalanceToFile(currentBalance *float64) {
	balanceString := fmt.Sprint(*currentBalance)
	os.WriteFile("XYZ Balance.txt", []byte(balanceString), 0644)
}
