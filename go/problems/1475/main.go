package main

func finalPrices(prices []int) []int {
	n := len(prices)
	result := make([]int, n)
	copy(result, prices)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= prices[i] {
			result[stack[len(stack)-1]] -= prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}
