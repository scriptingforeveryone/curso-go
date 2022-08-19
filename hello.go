package main

import "fmt"

func main() {
	a := 10

	var ponteiro *int = &a

	*ponteiro = 50

	b := &a
	fmt.Println(*b)
}
