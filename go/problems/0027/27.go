// 27. Remove Element
package rm_elem

func removeElement(nums []int, val int) int {
	writeIndex := 0

	for _, num := range nums {
		if num != val {
			nums[writeIndex] = num
			writeIndex++
		}
	}

	return writeIndex
}
