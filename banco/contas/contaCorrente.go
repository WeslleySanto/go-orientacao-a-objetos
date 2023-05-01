package contas

import "go-orientacao-a-objetos/banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valueIsGreaterThanZero(valorDoSaque) && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	}
	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valueIsGreaterThanZero(valorDoDeposito) {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.saldo
	}
	return "Valor do depósito menor que zero.", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valueIsGreaterThanZero(valorDaTransferencia) && valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func valueIsGreaterThanZero(valor float64) bool {
	return valor > 0
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
