package atm

import (
	"errors"
	"fmt"
	"strconv"
)

func Withdrawal(amount string) (map[int]int, error) {
	err := validateAmount(amount)
	if err != nil {
		return nil, err
	}
	cash := make(map[int]int, 6)
	var banknotesAmount int
	denominations := denominations()
	banknotes := getBanknotes()
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
			}
		}
		if amountToInt != 0 {
			fmt.Println("Insufficient funds in the ATM. Input other amount.")
			return nil, errors.New("insufficient funds in the ATM")
		}
		changeBanknotesAmount(cash)
	} else {
		fmt.Println(err)
	}
	return cash, nil
}
