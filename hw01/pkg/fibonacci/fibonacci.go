package fibonacci

import "fmt"

// Calc вычисление числа Фибоначчи по номеру
func Calc(number int) (result int, err error) {
	if number > 20 {
		return number, fmt.Errorf("- аргумент не может быть больше 20")
	} else if number == 1 {
		return 0, nil
	}

	firstNumber, result, tempNumber := 0, 1, 0

	for i := 2; i < number; i++ {
		tempNumber = result
		result += firstNumber
		firstNumber = tempNumber
	}
	return result, nil
}
