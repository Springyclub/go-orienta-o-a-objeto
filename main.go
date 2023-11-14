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

	countGui := CurrentAccount{holder: "Gui", numberAccount: 12, numberAgency: 21, balance: 200}
	countYas := CurrentAccount{holder: "Yas", numberAccount: 12, numberAgency: 21, balance: 200}

	status := countYas.Transfer(0, &countGui)

	fmt.Println(status)

	fmt.Println(countYas)
	fmt.Println(countGui)

}

func (c *CurrentAccount) Transfer(valueTransfer float64, countDestiny *CurrentAccount) bool {
	if valueTransfer <= c.balance && valueTransfer > 0 {
		c.balance -= valueTransfer
		countDestiny.Deposit(valueTransfer)
		return true
	}
	return false
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
