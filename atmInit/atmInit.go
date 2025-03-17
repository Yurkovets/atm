package atmInit

import (
	"fmt"
	"sort"
)

func AtmInit() {
	banknotes[5] = 20
	banknotes[10] = 20
	banknotes[100] = 15
	banknotes[200] = 15
	banknotes[500] = 10
	banknotes[1000] = 5

	fmt.Println("Enter the withdrawal amount: ")
}

func makeSliceOfAvailableDenominations(banknotes map[int]int) []int {
	var denominations = make([]int, 0)
	for denomination, amount := range banknotes {
		if amount != 0 {
			denominations = append(denominations, denomination)
		}
	}
	sort.Slice(denominations, func(i, j int) bool {
		return denominations[i] < denominations[j]
	})
	return denominations
}
