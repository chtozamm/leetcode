// 80. Remove Duplicates from Sorted Array II
package rm_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIdx := 1
	dupCount := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[writeIdx-1] {
			nums[writeIdx] = nums[i]
			writeIdx++
			dupCount = 1
		} else if dupCount < 2 {
			nums[writeIdx] = nums[i]
			writeIdx++
			dupCount++
		}
	}

	return writeIdx
}
