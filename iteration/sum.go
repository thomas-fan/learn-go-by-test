package iteration

// Sum add arrays
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll add to slice
func SumAll(numbersToSum ...[]int) (sums [] int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

// SumTails add tail
func SumTails(numbersToSumTail ...[]int) (sums [] int) {
	for _, numbers := range numbersToSumTail {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return
}