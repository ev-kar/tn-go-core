package main

import (
	"fmt"

	"../pkg/fibonacci"
)

func main() {
	buf := []int{1, 2, 3, 4, 5, 6, 7, 8, 13, 18, 21}

	for _, v := range buf {
		result := fibonacci.Calc(v)

		if v > 20 {
			fmt.Println(v, "- аргумент не может быть больше 20")
			return
		}

		fmt.Println("Число Фибоначчи №", v, "равно", result)
	}
}
