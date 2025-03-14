package atmInit

var withdrawalAmount string = ""
var banknotes = make(map[int]int, 6)
var denominations = make([]int, 6)

func GetBanknotes() map[int]int {
	return banknotes
}

func GetDenominations() []int {
	denominations = makeSliceOfAvailableDenominations(banknotes)
	return denominations
}

func GetWithdrawalAmount() *string {
	return &withdrawalAmount
}
