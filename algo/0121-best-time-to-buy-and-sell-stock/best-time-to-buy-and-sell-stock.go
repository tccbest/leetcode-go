package _121_best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	plen := len(prices)
	if plen == 0 {
		return 0
	}

	ret := 0
	minPrice := math.MaxInt64
	for i := 0; i < plen; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] - minPrice > ret {
			ret = prices[i] - minPrice
		}
	}

	return ret
}