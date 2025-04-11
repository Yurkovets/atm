package atm

type Atm struct {
	Banknotes map[int]int
}

func NewAtm(banknotes map[int]int) *Atm {
	return &Atm{
		Banknotes: banknotes,
	}
}
