package atm

import (
	"errors"
	"math"
)

func validateAmount(withdrawalAmount int) error {
	if withdrawalAmount < 0 {
		return errors.New("Please, check your input is correct. Negative numbers are not allowed.")
	}

	if withdrawalAmount > 10000 {
		return errors.New("Maximum cash withdrawal amount is 10000 dollars")
	}

	if math.Mod(float64(withdrawalAmount), 5) != 0 {
		return errors.New("Incorrect amount. Available denominations:" + printAvailableDenominations(denominations()))
	}

	return nil
}
