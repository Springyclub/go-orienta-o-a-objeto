package main

import "fmt"

type CurrentAccount struct {
	holder        string
	numberAccount int
	numberAgency  int
	balance       float64
}

func main() {

	countGui := CurrentAccount{}

	countGui.holder = "Gui"
	countGui.balance = 500

	withdrawalAmount := 200.

	fmt.Println(countGui.withdraw(withdrawalAmount))
	fmt.Println(countGui)

	status, valor := countGui.Deposit(500)

	fmt.Println(status, valor)

	fmt.Println(countGui)

}

func (c *CurrentAccount) Deposit(valueDeposit float64) (string, float64) {
	if valueDeposit < 1 {
		return "Valor inválido", 0
	}
	c.balance += valueDeposit

	return "Deposito resgatado com sucesso", c.balance
}
func (c *CurrentAccount) withdraw(withdrawalAmount float64) string {

	if withdrawalAmount > 0 {
		return "Valor para saque inválido"
	}

	canWithDraw := withdrawalAmount <= c.balance

	if canWithDraw {
		c.balance -= withdrawalAmount
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}
