// 912. Sort an Array
package sort_array

// sortArray sorts a slice of integers using **Counting Sort** algorithm.
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Count frequencies of numbers in `nums`. Note that we
	// subtract `min` to normalize the indeces to start from 0.
	count := make([]int, max-min+1)
	for _, num := range nums {
		count[num-min]++
	}

	// Write sorted numbers back to the input array.
	index := 0
	for i := range count {
		for count[i] > 0 {
			nums[index] = i + min
			index++
			count[i]--
		}
	}

	return nums
}
