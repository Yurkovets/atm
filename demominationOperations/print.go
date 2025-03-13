package demominationOperations

import (
	"fmt"
	"slices"
)

func PrintAvailableDenominations(denominations []int) {
	slices.Sort(denominations)
	for _, denomination := range denominations {
		fmt.Println(denomination)
	}
}
