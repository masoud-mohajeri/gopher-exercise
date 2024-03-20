package leetcode

import (
	"fmt"
	"time"
)

func ClimbStairs() {

	var climb func(int, map[int]int) int
	cache := make(map[int]int)
	cache[1] = 1
	cache[2] = 2

	climb = func(s int, c map[int]int) int {
		if s <= 0 {
			return 0
		}

		if val, ok := c[s]; ok {
			return val
		}

		if s == 1 {
			return 1
		}

		if s == 2 {
			return 2
		}
		calc := climb(s-1, c) + climb(s-2, c)
		c[s] = calc
		return calc
	}

	for i := 0; i < 45; i++ {
		start := time.Now()
		variations := climb(i, cache)
		elapsed := time.Since(start) / 1_000

		fmt.Printf("%d ns - for %d stairs: %d \n", elapsed, i, variations)
	}
}
