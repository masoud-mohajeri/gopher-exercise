package leetcode

import (
	"fmt"
)

func CoinChange() {
	coins := []int{1, 2, 5}

	var calc func([]int, int, map[int]int) int
	cache := make(map[int]int)

	calc = func(coins []int, rem int, cache map[int]int) int {
		if rem < 0 {
			return -1
		}
		if rem == 0 {
			return 0
		}
		if cache[rem] != 0 {
			return cache[rem]
		}

		bestResult := rem + 1

		for _, coin := range coins {
			res := calc(coins, rem-coin, cache)
			if res >= 0 && res < bestResult {
				bestResult = res
			}
		}
		if bestResult == rem+1 {
			return -1
		}
		cache[rem] = bestResult + 1

		return cache[rem]
	}

	fmt.Printf("result: %d \n", calc(coins, 100, cache))
	fmt.Println("cache:", cache)

}
