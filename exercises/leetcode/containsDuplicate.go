package leetcode

import (
	"fmt"
)

func ContainsDuplicate() bool {
	// nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	nums := []int{1, 2, 3, 1, 4, 5}
	numsMap := make(map[int]int)

	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			fmt.Println(true)
			return true
		}
		numsMap[num] = 1
	}

	fmt.Println(false)
	return false

	// using sort package !
	// no map

	// if len(nums) <= 1 {
	// 	return false
	// }
	// sort.Ints(nums)

	// // Time: O(n)
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] == nums[i+1] {
	// 		fmt.Println(true)

	// 		return true
	// 	}
	// }
	// fmt.Println(false)
	// return false

}
