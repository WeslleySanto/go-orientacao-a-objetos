package main

import (
	"fmt"
	"go-orientacao-a-objetos/banco/contas"
)

func PagarBoleto(conta conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type conta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 1000)

	fmt.Println(contaDaLuisa.ObterSaldo())
}
