package main

import (
	"fmt"

	"../pkg/fibonacci"
)

func main() {
	buf := []int{1, 2, 3, 4, 5, 6, 7, 8, 13, 18, 21}

	for _, v := range buf {
		result, err := fibonacci.Calc(v)

		if err != nil {
			fmt.Println(result, err)
			return
		}

		fmt.Println("Число Фибоначчи №", v, "равно", result)
	}
}
