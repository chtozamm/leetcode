package main

func minOperations(boxes string) []int {
	n := len(boxes)
	ans := make([]int, n)

	ballsToLeft := 0
	movesToLeft := 0
	for i := 0; i < n; i++ {
		ans[i] = movesToLeft
		if boxes[i] == '1' {
			ballsToLeft++
		}
		movesToLeft += ballsToLeft
	}

	ballsToRight := 0
	movesToRight := 0
	for i := n - 1; i >= 0; i-- {
		ans[i] += movesToRight
		if boxes[i] == '1' {
			ballsToRight++
		}
		movesToRight += ballsToRight
	}

	return ans
}
