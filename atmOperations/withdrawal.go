package atmOperations

import (
	"fmt"
	"strconv"
)

func CashWithdrawal(amount string, denominations []int, banknotes map[int]int) map[int]int {
	cash := make(map[int]int, 6)
	amountToInt, err := strconv.Atoi(amount)
	fmt.Println("User Input: ", amountToInt)
	if err == nil {
		for i := len(denominations) - 1; i >= 0; i-- {
			if amountToInt/denominations[i] != 0 {
				if banknotes[denominations[i]] >= amountToInt/denominations[i] {
					cash[denominations[i]] = amountToInt / denominations[i]
					amountToInt = amountToInt - denominations[i]*cash[denominations[i]]
				} else {
					cash[denominations[i]] = banknotes[denominations[i]]
					amountToInt = amountToInt - denominations[i]*cash[denominations[i]]
				}
			}
		}
		if amountToInt == 0 {
			return cash
		} else {
			fmt.Println("Insufficient funds in the ATM!")
			return cash
		}

	} else {
		fmt.Println(err)
		return cash
	}

}
