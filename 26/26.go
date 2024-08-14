// 26. Remove Duplicates from Sorted Array
package rm_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex-1] {
			nums[uniqueIndex] = nums[i]
			uniqueIndex++
		}
	}

	return uniqueIndex
}
