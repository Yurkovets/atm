package main

import (
	"fmt"

	"com.atm/atmInit"
	"com.atm/validations"
)

func main() {
	for {
		atmInit.AtmInit()
		fmt.Scanln(atmInit.GetWithdrawalAmount())
		validations.UserInputValidation(*atmInit.GetWithdrawalAmount(), atmInit.GetDenominations())
	}
}
