package learning

import (
	"fmt"
	count "go-orienta-o-a-objeto/learning/count"
)

func pointer() {

	account := count.CurrentAccount{
		Holder:        "Guilherme",
		NumberAccount: 589,
		NumberAgency:  123456,
		salance:       125.50,
	}
	account2 := count.CurrentAccount{
		Holder:        "Guilherme",
		NumberAccount: 589,
		NumberAgency:  123456,
		salance:       125.50,
	}
	fmt.Println("Count of "+account.Holder+"\nNumber agency:"+
		"", account.NumberAgency, "\nNumber account: ", account.NumberAgency, ""+
		"\nBalance:", account.salance)

	fmt.Println("Count of "+account2.Holder+"\nNumber agency:"+
		"", account2.NumberAgency, "\nNumber account: ", account2.NumberAgency, ""+
		"\nBalance:", account2.salance)

	fmt.Println(account == account2)
	fmt.Println("\n\n\n\n\n\n")
	var countYasmin *count.CurrentAccount
	countYasmin = new(count.CurrentAccount)
	countYasmin.Holder = "Yasmin"
	countYasmin.salance = 800

	fmt.Println(*countYasmin, "Ponteiro")
	fmt.Println(countYasmin, "Sem ponteiro")

	var countYasmin2 *count.CurrentAccount
	countYasmin2 = new(count.CurrentAccount)
	countYasmin2.Holder = "Yasmin"
	countYasmin2.salance = 800

	fmt.Println(*countYasmin2, "Ponteiro")
	fmt.Println(countYasmin2, "Sem ponteiro\n\n")

	fmt.Println(&countYasmin)
	fmt.Println(&countYasmin2)
	fmt.Println(countYasmin2, "Sem ponteiro")

	fmt.Println(*countYasmin == *countYasmin2)
}
