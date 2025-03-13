package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func main() {

	banknotes := map[int]int{5: 1, 10: 2, 100: 3, 200: 25, 500: 25, 1000: 10}
	denominations := makeSliceOfAvailableDenominations(banknotes)
	for {
		var withdrawalAmount string
		fmt.Println("Enter the withdrawal amount: ")
		fmt.Scanln(&withdrawalAmount)

		amount, err := strconv.Atoi(withdrawalAmount)
		if err != nil {
			fmt.Println("Input error. Plese, input a number value.")
		}
		if math.Mod(float64(amount), 5) != 0 {
			fmt.Println("Incorrect amount. Available denominations:")
			printAvailableDenominations(denominations)
		}
	}
}

func printAvailableDenominations(denominations []int) {
	slices.Sort(denominations)
	for _, denomination := range denominations {
		fmt.Println(denomination)
	}
}

func makeSliceOfAvailableDenominations(banknotes map[int]int) []int {
	var denominations = make([]int, 0)
	for denomination, amount := range banknotes {
		if amount != 0 {
			denominations = append(denominations, denomination)
		}
	}
	return denominations
}

// func cashWithdrawal(amount int) map[int]int {

// }
