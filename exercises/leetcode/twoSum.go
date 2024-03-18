package leetcode

import "fmt"

func TwoSum() {
	nums := []int{1, 2, 3, 4, 5}
	target := 7

	// brute force
	bfAnswers := [][]int{}
	for _, numi := range nums {
		for _, numj := range nums {
			sum := numi + numj
			if sum == target {
				bfAnswers = append(bfAnswers, []int{numj, numi})
			}
		}
	}

	// using maps
	numMap := make(map[int]int)
	mapAnswers := [][]int{}
	for _, numi := range nums {
		defect := target - numi

		if _, ok := numMap[defect]; ok {
			mapAnswers = append(mapAnswers, []int{defect, numi})
		}
		numMap[numi] = defect
	}

	fmt.Println("bfAnswers:", bfAnswers)
	fmt.Println("mapAnswers:", mapAnswers)
}
