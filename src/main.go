package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {

	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(200))

	fmt.Println(contaDaSilvia.Depositar(2000))

	fmt.Println(contaDaSilvia.saldo)
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "O valor do deposito é menor que 0.", c.saldo
	}
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo indisponivel."
	}

}
