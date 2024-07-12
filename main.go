package main

import (
	"fmt"
	"example.com/meuprojeto/contas"
	//"example.com/meuprojeto/clientes"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64)  {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64)string
}


func main() {
	contaNavarro := contas.ContaPoupanca{}
	contaNavarro.Depositar(100)
	PagarBoleto(&contaNavarro,60)

	contaJoao := contas.ContaCorrente{}
	contaJoao.Depositar(299)
	PagarBoleto(&contaJoao,157)


	fmt.Println(contaJoao)
	fmt.Println(contaNavarro)

}
