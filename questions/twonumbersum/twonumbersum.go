package main

func TwoNumberSum(array []int, targetSum int) (answer []int) {
	answer = make([]int, 0, 2)
	size := len(array)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			if array[i]+array[j] == targetSum {
				answer = []int{array[i], array[j]}
				return
			}
		}
	}
	return
}
