package main

import (
	"fmt"
	"strconv"

	"github.com/Yurkovets/atm/atm"
)

func main() {
	var withdrawalAmount string = ""
	for {
		fmt.Println("Enter the withdrawal amount: ")
		fmt.Scanln(&withdrawalAmount)

		withdrawalAmountToInt, err := strconv.Atoi(withdrawalAmount)
		if err != nil {
			fmt.Println("Please, check your input is correct. Only number values are valid.")
			continue
		}

		cash, err := atm.Withdrawal(withdrawalAmountToInt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(cash)
	}
}
