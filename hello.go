package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carro{
		Nome: "BMW",
	}

	carro.andar()

	resultado := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}

		return func() int {
			return res * res
		}

	}
	// resultado := somaTudo(10, 20, 40)

	fmt.Println(resultado(54, 54, 54, 54)())

}
