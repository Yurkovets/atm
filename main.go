package main

import (
	"fmt"

	"com.atm/atmInit"
	"com.atm/atmOperations"
	"com.atm/validations"
)

func main() {

	atmInit.AtmInit()
	var withdrawalAmount string = ""

	for {
		banknotes := atmInit.GetBanknotes()
		denominations := atmInit.GetDenominations()

		fmt.Println("Enter the withdrawal amount: ")
		fmt.Scanln(&withdrawalAmount)

		validations.UserInputValidation(withdrawalAmount, denominations)
		cash := atmOperations.CashWithdrawal(withdrawalAmount, denominations, banknotes)
		fmt.Println(cash)
	}
}
