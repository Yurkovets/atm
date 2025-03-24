package atm

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func validateAmount(userInput string) error {
	withdrawalAmountToInt, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Please, check your input is correct. Only number values are valid.")
		return errors.New("string value provided")
	}

	if math.Mod(float64(withdrawalAmountToInt), 5) != 0 {
		fmt.Println("Incorrect amount. Available denominations:")
		printAvailableDenominations(denominations())
		return errors.New("incorrect amount")
	}

	if withdrawalAmountToInt > 10000 {
		fmt.Println("Maximum cash withdrawal amount is 10000 dollars")
		return errors.New("too much amount")
	}

	if withdrawalAmountToInt < 0 {
		fmt.Println("Please, check your input is correct. Negative numbers are not allowed.")
		return errors.New("negative amount")
	}

	return nil
}
