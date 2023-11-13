package learning

import "fmt"

type CurrentAccount struct {
	holder        string
	numberAccount int
	numberAgency  int
	balance       float64
}

func pointer() {

	account := CurrentAccount{
		holder:        "Guilherme",
		numberAccount: 589,
		numberAgency:  123456,
		balance:       125.50,
	}
	account2 := CurrentAccount{
		holder:        "Guilherme",
		numberAccount: 589,
		numberAgency:  123456,
		balance:       125.50,
	}
	fmt.Println("Count of "+account.holder+"\nNumber agency:"+
		"", account.numberAgency, "\nNumber account: ", account.numberAgency, ""+
		"\nBalance:", account.balance)

	fmt.Println("Count of "+account2.holder+"\nNumber agency:"+
		"", account2.numberAgency, "\nNumber account: ", account2.numberAgency, ""+
		"\nBalance:", account2.balance)

	fmt.Println(account == account2)
	fmt.Println("\n\n\n\n\n\n")
	var countYasmin *CurrentAccount
	countYasmin = new(CurrentAccount)
	countYasmin.holder = "Yasmin"
	countYasmin.balance = 800

	fmt.Println(*countYasmin, "Ponteiro")
	fmt.Println(countYasmin, "Sem ponteiro")

	var countYasmin2 *CurrentAccount
	countYasmin2 = new(CurrentAccount)
	countYasmin2.holder = "Yasmin"
	countYasmin2.balance = 800

	fmt.Println(*countYasmin2, "Ponteiro")
	fmt.Println(countYasmin2, "Sem ponteiro\n\n")

	fmt.Println(&countYasmin)
	fmt.Println(&countYasmin2)
	fmt.Println(countYasmin2, "Sem ponteiro")

	fmt.Println(*countYasmin == *countYasmin2)
}
