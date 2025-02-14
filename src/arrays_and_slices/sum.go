package main

func Sum(numbers [4]int) int {
	var total int
	for _, number := range numbers {
		total += number

	}
	return total
} 