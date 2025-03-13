package main

import (
	"fmt"

	"com.atm/atmInit"
	"com.atm/validations"
)

func main() {

	atmInit.AtmInit(atmInit.Banknotes)
	denominations := atmInit.MakeSliceOfAvailableDenominations(atmInit.Banknotes)
	for {
		fmt.Println("Enter the withdrawal amount: ")
		fmt.Scanln(&atmInit.WithdrawalAmount)
		validations.UserInputValidation(atmInit.WithdrawalAmount, denominations)
	}
}

// func cashWithdrawal(amount int) map[int]int {

// }
