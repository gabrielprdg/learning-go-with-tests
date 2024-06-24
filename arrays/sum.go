package arrays

func Sum(numbers []int) int {
	var sum int = 0

	for _, nb := range numbers {
		sum += nb
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sumAll []int
	for _, numbers := range numbersToSum {
		sumAll = append(sumAll, Sum(numbers))
	}

	return sumAll
}

func SumAllTails(numbers ...[]int) []int {
	var tails []int
	for _, number := range numbers {
		if len(number) == 0 {
			tails = append(tails, Sum(number))
		} else {
			number = number[1:]
			tails = append(tails, Sum(number))
		}
	}
	return tails
}
