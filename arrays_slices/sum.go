package arraysslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sum []int) {
	qtdeNumber := len(numbersToSum)
	sum = make([]int, qtdeNumber)

	for i, numeros := range numbersToSum {
		sum[i] = Sum(numeros)
	}
	return
}
