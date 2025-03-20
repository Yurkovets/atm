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

		err := validations.UserInputValidation(withdrawalAmount, denominations)
		if err != nil {
			continue
		}
		cash, err := atmOperations.CashWithdrawal(withdrawalAmount, denominations, banknotes)
		if err != nil {
			continue
		}
		fmt.Println(cash)
	}
}
