package validations

import (
	"fmt"
	"math"
	"strconv"

	"com.atm/atmOperations"
)

func UserInputValidation(userInput string, denominations []int) {
	withdrawalAmountToInt, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Input error. Plese, input a number value.")
	} else if math.Mod(float64(withdrawalAmountToInt), 5) != 0 {
		fmt.Println("Incorrect amount. Available denominations:")
		atmOperations.PrintAvailableDenominations(denominations)
	}
}
