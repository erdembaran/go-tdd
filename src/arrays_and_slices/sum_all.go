package main

func SumAll(slice1, slice2 []int) []int {
	sum1, sum2 := 0, 0
	
	for _, v := range slice1 {
		sum1 += v
	}
	for _, v := range slice2 {
		sum2 += v
	}

	return []int{sum1, sum2}
}