// 15. 3Sum
package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	n := len(nums)
	for i := 0; i < n-2; i++ {
		// Skip duplicate values for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Early exit if the current number is greater than 0
		if nums[i] > 0 {
			break
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicate values for the second number
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// Skip duplicate values for the third number
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
