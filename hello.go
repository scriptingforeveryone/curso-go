package main

import "fmt"

func main() {
	a := 10

	var ponteiro *int = &a
	fmt.Println(*ponteiro)
}
