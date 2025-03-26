package main

import (
	"fmt"
	"strconv"

	"github.com/Yurkovets/atm/atm"
)

func main() {
	var input string = ""
	for {
		fmt.Println("Enter the withdrawal amount: ")
		fmt.Scanln(&input)

		withdrawalAmount, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please, check your input is correct. Only number values are valid.")
			continue
		}

		cash, err := atm.Withdrawal(withdrawalAmount)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(cash)
	}
}
