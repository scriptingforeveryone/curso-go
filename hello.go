package main

import (
	"errors"
	"fmt"
)

func main() {
	res, _ := soma(7, 2)
	fmt.Println(res)
}

func soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}
	return res, nil
}
