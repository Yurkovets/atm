package atm

import (
	"slices"
	"strconv"
)

func printAvailableDenominations(denominations []int) string {
	slices.Sort(denominations)
	var printedDenominations string
	for _, denomination := range denominations {
		printedDenominations += "\n" + strconv.Itoa(denomination)
	}

	return printedDenominations
}
