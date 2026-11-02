package arraysslices

import "math"

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

func FindMaxNumber(numbers []int) int {
	if (len(numbers) == 0) {
		return 0;
	}

	maxNumber := numbers[0]

	// for _, number := range numbers {
	// 	if maxNumber < number {
	// 		maxNumber = number
	// 	}
	// }

	for i := 1; i < len(numbers); i++ {
		if maxNumber < numbers[i] {
			maxNumber = numbers[i]
		}
	}

	return maxNumber
}

func SplitMiddle(slice []int) [2][]int {
	if len(slice) == 0 {
		return [2][]int {}
	}

	middle := int(math.Trunc(float64(len(slice) / 2)))
	left := slice[:middle]
	rigth := slice[middle:]

	return [2][]int{
		left,
		rigth,
	}
}