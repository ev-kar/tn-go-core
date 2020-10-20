package fibonacci

// Calc вычисление числа Фибоначчи по номеру
func Calc(number int) (result int) {
	if number == 1 {
		return 0
	}

	firstNumber, result := 0, 1

	for i := 2; i < number; i++ {
		result, firstNumber = firstNumber, result
		result += firstNumber
	}
	return result
}
