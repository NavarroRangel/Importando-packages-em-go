package main

import (
	"fmt"
	"example.com/meuprojeto/contas"
)



func main() {
	contaNavarro := contas.ContaCorrente {Titular: "navarro", Saldo: 300}
	contaJoao := contas.ContaCorrente {Titular: "joao", Saldo: 500}

	status := contaNavarro.Transferir(-200,&contaJoao)
	
	
	fmt.Println(status)
	fmt.Println(contaNavarro)
	fmt.Println(contaJoao)

	;


	
}
