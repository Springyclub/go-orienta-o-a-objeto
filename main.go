package main

import (
	"fmt"
)

type CurrentAccount struct {
	holder        string
	numberAccount int
	numberAgency  int
	balance       float64
}

func main() {

	account := CurrentAccount{
		holder:        "Guilherme",
		numberAccount: 589,
		numberAgency:  123456,
		balance:       125.50,
	}
	fmt.Println("Count of "+account.holder+"\nNumber agency:"+
		"", account.numberAgency, "\nNumber account: ", account.numberAgency, ""+
		"\nBalance:", account.balance)

	var countYasmin *CurrentAccount
	countYasmin = new(CurrentAccount)
	countYasmin.holder = "Yasmin"
	countYasmin.balance = 800

	fmt.Println(*countYasmin, "Ponteiro")
	fmt.Println(countYasmin, "Sem ponteiro")
}
