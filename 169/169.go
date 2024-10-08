// 169. Majority Element
package major_elem

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
		}

		if count == 0 {
			candidate = nums[i]
			count = 1
		}
	}

	return candidate
}
