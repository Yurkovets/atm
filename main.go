package main

import (
	"fmt"
	"strconv"

	atm "github.com/Yurkovets/atm/atm"
)

func main() {

	var banknotes = map[int]int{
		5:    20,
		10:   20,
		100:  15,
		200:  15,
		500:  10,
		1000: 5,
	}
	var input string = ""
	atm := atm.NewAtm(banknotes)
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
