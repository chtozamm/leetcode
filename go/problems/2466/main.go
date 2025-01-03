package main

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1

	mod := 1_000_000_007

	for end := 1; end <= high; end++ {
		if end >= zero {
			dp[end] = (dp[end] + dp[end-zero]) % mod
		}
		if end >= one {
			dp[end] = (dp[end] + dp[end-one]) % mod
		}
	}

	ans := 0
	for i := low; i <= high; i++ {
		ans = (ans + dp[i]) % mod
	}
	return ans
}
