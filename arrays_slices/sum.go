package arraysslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, numeros := range numbersToSum {
		sum = append(sum, Sum(numeros))
	}
	return sum
}

func SumAllLastRest(numberstoSum ...[]int) []int {
	var sum []int
	for _, numeros := range numberstoSum {
		if len(numeros) == 0 {
			sum = append(sum, 0)
		} else {
			final := numeros[1:]
			sum = append(sum, Sum(final))
		}
	}
	return sum
}
