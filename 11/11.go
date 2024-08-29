// 11. Container With Most Water
package water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var maxContainerArea int

	for left < right {
		area := (right - left) * min(height[right], height[left])
		maxContainerArea = max(maxContainerArea, area)

		// Move the pointer for the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxContainerArea
}
