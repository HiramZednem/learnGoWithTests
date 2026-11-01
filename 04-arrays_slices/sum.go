package arraysslices

func Sum(numbers []int) int {
	var sum int;
	for _, number := range numbers {
		sum += number;
	}
	return sum;
}

// Implementation worrying about memory and capacity
// func SumAll(numbersToSum ...[]int) []int {
// 	lenghtOfNumbers := len(numbersToSum)
// 	sums := make([]int, lenghtOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
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
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}