package main

import (
	"fmt"
	"example.com/meuprojeto/contas"
	//"example.com/meuprojeto/clientes"
)



func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)
	fmt.Println(contaExemplo.ObterSaldo())
}
