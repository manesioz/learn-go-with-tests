package main

//Sum sums the elements in slice
func Sum(arrays ...[]int) []int {
	result := []int{}
	for _, arr := range arrays {
		count := 0
		for _, elem := range arr {
			count += elem
		}
		result = append(result, count)
	}
	return result
}
