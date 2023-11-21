package arrays

func Sum(numbers []int) int {
	var sum int

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumber := len(numbersToSum)
	// sums := make([]int, lengthOfNumber)
	var sums []int

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		// refactor: 使用append避免新增元素的時候超過容量，如超過會產生panic。
		sums = append(sums, Sum(numbers))

	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sumTails := numbers[1:]
			sums = append(sums, Sum(sumTails))
		}
	}

	return sums
}
