package contas

import "go-orientacao-a-objetos/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valueIsGreaterThanZero(valorDoSaque) && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	}
	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valueIsGreaterThanZero(valorDoDeposito) {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.saldo
	}
	return "Valor do depósito menor que zero.", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
