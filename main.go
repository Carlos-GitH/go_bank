package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	// 	// var titular string = "John"
	// 	// var numeroAgencia int = 123
	// 	// var numeroConta int = 123456
	// 	// var saldo float64 = 0.0
	// 	// println(titular)
	// 	// println(numeroAgencia)
	// 	// println(numeroConta)
	// 	// println(saldo)
	// 	conta1 := contas.ContaCorrente{Titular.: "John", NumeroAgencia: 123, NumeroConta: 123456, Saldo: 1000}
	// 	// fmt.Println(conta1)

	// 	var conta2 contas.ContaCorrente
	// 	conta2.Titular = "Jane"
	// 	conta2.NumeroAgencia = 123
	// 	conta2.NumeroConta = 123456
	// 	conta2.Saldo = 0.0
	// 	// fmt.Println(conta2)

	// 	fmt.Println(conta1.Sacar(117.23))

	// 	// status, valor := fmt.Println(conta1.Depositar(117.23))
	// 	// fmt.Println("status: ", status, " valor: ", valor)

	// 	status := conta1.Transferir(117.23, &conta2)
	// 	fmt.Println(status)
	// 	fmt.Println(conta1.Saldo)
	// 	fmt.Println(conta2.Saldo)
	//

	// conta1 := contas.ContaCorrente{}
	// conta1.Titular = clientes.Titular{
	// 	Nome:      "John",
	// 	CPF:       "123456",
	// 	Profissao: "Developer",
	// }
	// conta1.NumeroAgencia = 123
	// conta1.NumeroConta = 123456
	// conta1.Saldo = 1250.52
	// fmt.Println(conta1)

	cliente1 := clientes.Titular{
		Nome:      "John",
		CPF:       "123456",
		Profissao: "Developer",
	}

	conta1 := contas.ContaCorrente{
		Titular:       cliente1,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	conta1.Depositar(1050.0)
	fmt.Println(conta1)

	PagarBoleto(&conta1, 117.23)
	fmt.Println(conta1.ObterSaldo())

	cliente2 := clientes.Titular{
		Nome:      "Jane",
		CPF:       "123456",
		Profissao: "Designer",
	}
	conta2 := contas.ContaPoupanca{
		Titular:       cliente2,
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Operacao:      1,
	}
	conta2.Depositar(1921.23)
	fmt.Println(conta2.ObterSaldo())
	PagarBoleto(&conta2, 917.23)
	fmt.Println(conta2.ObterSaldo())
}
