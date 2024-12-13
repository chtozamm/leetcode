// 122. Best Time to Buy and Sell Stock II
package profit

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
