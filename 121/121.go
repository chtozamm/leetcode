// 121. Best Time to Buy and Sell Stock
package profit

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	minPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > profit {
			profit = price - minPrice
		}
	}

	return profit
}
