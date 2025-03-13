package atmInit

func AtmInit(banknotes map[int]int) {
	banknotes[5] = 20
	banknotes[10] = 20
	banknotes[100] = 15
	banknotes[200] = 15
	banknotes[500] = 10
	banknotes[1000] = 5
}

func MakeSliceOfAvailableDenominations(banknotes map[int]int) []int {
	var denominations = make([]int, 0)
	for denomination, amount := range banknotes {
		if amount != 0 {
			denominations = append(denominations, denomination)
		}
	}
	return denominations
}
