package atm

import (
	"errors"
)

func Withdrawal(amount int) (map[int]int, error) {
	err := validateAmount(amount)
	if err != nil {
		return nil, err
	}

	denominations := denominations()
	banknotes := getBanknotes()
	cash := make(map[int]int)

	for i := len(denominations) - 1; i >= 0; i-- {
		denomination := denominations[i]
		amountToDenom := amount / denomination
		if amountToDenom == 0 {
			continue
		}

		banknotesAmount := banknotes[denomination]
		if banknotesAmount >= amountToDenom {
			cash[denomination] = amountToDenom
		} else {
			cash[denomination] = banknotesAmount
		}
		amount = amount - denomination*cash[denomination]
	}

	if amount != 0 {
		return nil, errors.New("Insufficient funds in the ATM. Input other amount.")
	}

	changeBanknotesAmount(cash)

	return cash, nil
}
