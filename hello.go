package main

import "fmt"

func main() {
	a := 10

	var ponteiro *int = &a

	*ponteiro = 50
	fmt.Println(*ponteiro)
}
