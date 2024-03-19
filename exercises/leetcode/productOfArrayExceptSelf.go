package leetcode

import "fmt"

func ProductOfArrayExceptSelf() {
	nums := []int{-1, 1, 0, -3, -2}
	zerosCount := 0
	product := 1
	answer := make([]int, len(nums))

	for _, num := range nums {
		if num == 0 {
			zerosCount++
		} else {
			product = product * num
		}
	}
	//	if zerosCount == 1 { zero values!!! }
	if zerosCount == 1 {
		for i, num := range nums {
			if num == 0 {
				answer[i] = product
			} else {
				answer[i] = 0
			}
		}
	}
	if zerosCount == 0 {
		for i, num := range nums {
			answer[i] = product / num
		}
	}

	fmt.Println("answer:", answer)
}
