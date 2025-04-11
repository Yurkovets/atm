package atm

type atm struct {
	banknotes map[int]int
}

func newAtm(banknotes map[int]int) *atm {
	return &atm{
		banknotes: banknotes,
	}
}

func (a *atm) getBanknotes() map[int]int {
	return a.banknotes
}
