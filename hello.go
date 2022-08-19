package main

import "fmt"

func main() {

	resultado := somaTudo(10, 20, 40)

	fmt.Println(resultado)

}

func soma(a int, b int) (result int) {

	result = a + b
	return

}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}
	return resultado
}
