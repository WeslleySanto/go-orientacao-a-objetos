package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valueIsGreaterThanZero(valorDoSaque) && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	}
	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valueIsGreaterThanZero(valorDoDeposito) {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.Saldo
	}
	return "Valor do depósito menor que zero.", c.Saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valueIsGreaterThanZero(valorDaTransferencia) && valorDaTransferencia < c.Saldo {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func valueIsGreaterThanZero(valor float64) bool {
	return valor > 0
}
