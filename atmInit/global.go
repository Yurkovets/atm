package atmInit

var banknotes = make(map[int]int, 6)
var denominations = make([]int, 6)

func GetBanknotes() map[int]int {
	return banknotes
}

func ChangeBanknotesAmount(denomination int, amount int) {
	banknotes[denomination] = amount
}

func GetDenominations() []int {
	denominations = makeSliceOfAvailableDenominations(banknotes)
	return denominations
}
