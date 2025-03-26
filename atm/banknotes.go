package atm

import (
	"slices"
	"sort"
	"strconv"
)

var banknotes = map[int]int{
	5:    20,
	10:   20,
	100:  15,
	200:  15,
	500:  10,
	1000: 5,
}

func getBanknotes() map[int]int {
	return banknotes
}

func changeBanknotesAmount(cashIssued map[int]int) {
	for denomination, amount := range cashIssued {
		banknotes[denomination] -= amount
	}
}

func denominations() []int {
	var denominations []int
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

func getDenominationsString(denominations []int) string {
	slices.Sort(denominations)
	var printedDenominations string
	for _, denomination := range denominations {
		printedDenominations += "\n" + strconv.Itoa(denomination)
	}

	return printedDenominations
}
