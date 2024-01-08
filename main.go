package main

import (
	"fmt"
	clients "go-orienta-o-a-objeto/learning/clients"
	count "go-orienta-o-a-objeto/learning/count"
)

func PayBill(count verifyCount, valueBill float64) {
	count.Withdraw(valueBill)
}

type verifyCount interface {
	Withdraw(value float64) string
}

func main() {

	countGui := count.CountSavings{}

	countGui.Deposit(100)

	fmt.Println(countGui.GetBalance())
	PayBill(&countGui, 60)
	fmt.Println(countGui.GetBalance())

	fmt.Println(countGui.GetBalance())
	clientGui := clients.Holder{
		Name:       "Guigas",
		Document:   "13853314961",
		Profession: "Trainee developher",
	}

	countEnzo := count.CountSavings{
		Holder:        clientGui,
		NumberAccount: 12,
		NumberAgency:  12,
		Operation:     12,
	}

	countEnzo.Deposit(500)
	PayBill(&countEnzo, 400)

}
