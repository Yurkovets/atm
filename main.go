package main

import (
	"fmt"
	"math"
)

func main() {

	banknotes := map[int]int{5: 200, 10: 100, 100: 50, 200: 25, 500: 25, 1000: 10}

	for {
		fmt.Println("Choose option:")
		fmt.Println("1 - cash withdrawal")
		fmt.Println("2 - exit")

		var option int
		var withdrawalAmount int
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Println("Enter the withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)
			if math.Mod(float64(withdrawalAmount), 5) != 0 {
				fmt.Println("Incorrect amount. Available denominations:")
				printAvailableDenominations(banknotes)
			}

		case 2:
			return
		default:
			fmt.Println("No such option exists")
			fmt.Println(option)
		}
	}
}

func printAvailableDenominations(banknotes map[int]int) {
	for denomination, amount := range banknotes {
		if amount > 0 {
			fmt.Println(denomination)
		}
	}
}
