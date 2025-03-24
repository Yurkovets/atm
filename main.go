package main

import (
	"fmt"

	"github.com/Yurkovets/atm/atm"
)

func main() {
	var withdrawalAmount string = ""
	for {
		fmt.Println("Enter the withdrawal amount: ")
		fmt.Scanln(&withdrawalAmount)

		cash, err := atm.Withdrawal(withdrawalAmount)
		if err != nil {
			continue
		}
		fmt.Println(cash)
	}
}
