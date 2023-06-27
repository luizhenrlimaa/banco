package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {

	contaDoLuiz := ContaCorrente{"Luiz", 589, 123456, 125.5}
	contaDaBruna := ContaCorrente{"Bruna", 789, 111222, 200}

	fmt.Println(contaDoLuiz)
	fmt.Println(contaDaBruna)
}
