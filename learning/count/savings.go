package learning

import "go-orienta-o-a-objeto/learning/clients"

type CountSavings struct {
	Holder                                 clients.Holder
	NumberAgency, NumberAccount, Operation int
	balance                                float64
}

func (c *CountSavings) Transfer(valueTransfer float64, countDestiny *CurrentAccount) bool {
	if valueTransfer <= c.balance && valueTransfer > 0 {
		c.balance -= valueTransfer
		countDestiny.Deposit(valueTransfer)
		return true
	}
	return false
}
func (c *CountSavings) Deposit(valueDeposit float64) (string, float64) {
	if valueDeposit < 1 {
		return "Valor inválido", 0
	}
	c.balance += valueDeposit

	return "Deposito resgatado com sucesso", c.balance
}

func (c *CountSavings) GetBalance() float64 {
	return c.balance
}

func (c *CountSavings) Withdraw(withdrawalAmount float64) string {

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
