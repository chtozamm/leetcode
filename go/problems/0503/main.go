package main

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	stack := make([]int, 0)
	for i := 0; i < n*2; i++ {
		currentIndex := i % n

		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[currentIndex] {
			ans[stack[len(stack)-1]] = nums[currentIndex]
			stack = stack[:len(stack)-1]
		}

		if i < n {
			stack = append(stack, currentIndex)
		}

	}

	return ans
}
