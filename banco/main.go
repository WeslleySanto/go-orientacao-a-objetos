package main

import (
	"fmt"
	"go-orientacao-a-objetos/banco/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	msg, saldo := contaExemplo.Depositar(100)
	fmt.Println(msg, saldo)

	// fmt.Println(contaExemplo.ObterSaldo())
}
