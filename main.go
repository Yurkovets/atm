package main

import (
	"fmt"

	"com.atm/atmInit"
	"com.atm/atmOperations"
	"com.atm/validations"
)

func main() {
	for {
		atmInit.AtmInit()
		fmt.Scanln(atmInit.GetWithdrawalAmount())
		validations.UserInputValidation(*atmInit.GetWithdrawalAmount(), atmInit.GetDenominations())
		cash := atmOperations.CashWithdrawal(*atmInit.GetWithdrawalAmount(), atmInit.GetDenominations(), atmInit.GetBanknotes())
		fmt.Println(cash)
	}
}
