package leetcode

import "fmt"

func MaxProfit() {
	prices := []int{7, 1, 5, 3, 12, 4}
	bfProfit := 0
	// brute force
	for i, pi := range prices {
		for j := i; j < len(prices); j++ {
			profiti := prices[j] - pi
			if profiti > bfProfit {
				bfProfit = profiti
			}

		}
	}
	// time series ?!
	minPrice := prices[0]
	tsProfit := 0

	for i := 0; i < len(prices); i++ {
		currentPrice := prices[i]

		if currentPrice < minPrice {
			minPrice = currentPrice
		} else if currentPrice-minPrice > tsProfit {
			tsProfit = currentPrice - minPrice
		}
	}

	fmt.Println("bfProfit:", bfProfit)
	fmt.Println("tsProfit:", tsProfit)

}
