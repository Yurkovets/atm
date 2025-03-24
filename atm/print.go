package atm

import (
	"fmt"
	"slices"
)

func printAvailableDenominations(denominations []int) {
	slices.Sort(denominations)
	for _, denomination := range denominations {
		fmt.Println(denomination)
	}
}
