package contas
import "example.com/meuprojeto/clientes"


type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia,NumeroConta,Operacao 		int
	saldo 									float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Impossível depositar", c.saldo
	}
}

func (c *ContaPoupanca)ObterSaldo()float64  {
	return c.saldo
}