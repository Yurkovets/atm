package atmOperations

import (
	"fmt"
	"strconv"

	"com.atm/atmInit"
)

func CashWithdrawal(amount string, denominations []int, banknotes map[int]int) map[int]int {
	cash := make(map[int]int, 6)
	var banknotesAmount int
	amountToInt, err := strconv.Atoi(amount)
	if err == nil {
		for i := len(denominations) - 1; i >= 0; i-- {
			banknotesAmount = banknotes[denominations[i]]
			if amountToInt/denominations[i] != 0 {
				if banknotesAmount >= amountToInt/denominations[i] {
					cash[denominations[i]] = amountToInt / denominations[i]
				} else {
					cash[denominations[i]] = banknotesAmount
				}
				amountToInt = amountToInt - denominations[i]*cash[denominations[i]]
				atmInit.ChangeBanknotesAmount(denominations[i], (banknotesAmount - cash[denominations[i]]))
			}
		}
		if amountToInt != 0 {
			fmt.Println("Insufficient funds in the ATM!")
		}

	} else {
		fmt.Println(err)
	}
	return cash
}
