package maxprofit

func maxProfit(prices []int) int {
	maxProfit := 0

	if len(prices) > 0 {
		lowest := prices[0]
		for i := 0; i < len(prices); i++ {
			price := prices[i]

			if lowest > price {
				lowest = price
			}

			profit := price - lowest

			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
