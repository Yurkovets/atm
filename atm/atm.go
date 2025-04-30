package atm

import (
	"errors"
	"math"
	"slices"
	"sort"
	"strconv"
)

type Atm struct {
	banknotes map[int]int
}

func NewAtm(banknotes map[int]int) *Atm {
	return &Atm{
		banknotes: banknotes,
	}
}

func (a *Atm) Withdrawal(amount int) (map[int]int, error) {
	err := a.validateAmount(amount)
	if err != nil {
		return nil, err
	}

	denominations := a.denominations()
	cash := make(map[int]int)

	for i := len(denominations) - 1; i >= 0; i-- {
		denomination := denominations[i]
		amountToDenom := amount / denomination
		if amountToDenom == 0 {
			continue
		}

		banknotesAmount := a.banknotes[denomination]
		if banknotesAmount >= amountToDenom {
			cash[denomination] = amountToDenom
		} else {
			cash[denomination] = banknotesAmount
		}
		amount = amount - denomination*cash[denomination]
	}

	if amount != 0 {
		return nil, errors.New("Insufficient funds in the ATM. Input other amount.")
	}

	a.changeBanknotesAmount(cash)

	return cash, nil
}

func (a *Atm) validateAmount(withdrawalAmount int) error {
	if withdrawalAmount < 0 {
		return errors.New("Please, check your input is correct. Negative numbers are not allowed.")
	}

	if withdrawalAmount > 10000 {
		return errors.New("Maximum cash withdrawal amount is 10000 dollars")
	}

	if math.Mod(float64(withdrawalAmount), 5) != 0 {
		return errors.New("Incorrect amount. Available denominations:" + a.getDenominationsString())
	}

	return nil
}

func (a *Atm) changeBanknotesAmount(cashIssued map[int]int) {
	for denomination, amount := range cashIssued {
		a.banknotes[denomination] -= amount
	}
}

func (a *Atm) denominations() []int {
	var denominations []int
	for denomination, amount := range a.banknotes {
		if amount != 0 {
			denominations = append(denominations, denomination)
		}
	}
	sort.Slice(denominations, func(i, j int) bool {
		return denominations[i] < denominations[j]
	})
	return denominations
}

func (a *Atm) getDenominationsString() string {
	denominations := a.denominations()
	slices.Sort(denominations)
	var printedDenominations string
	for _, denomination := range denominations {
		printedDenominations += "\n" + strconv.Itoa(denomination)
	}

	return printedDenominations
}
