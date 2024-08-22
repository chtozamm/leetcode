// 42. Trapping Rain Water
package rain_water

func trap(height []int) int {
	n := len(height)
	totalWaterTrapped := 0
	var stack []int

	for rightIndex := 0; rightIndex < n; rightIndex++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= height[rightIndex] {
			// Pop the stack
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				// Calculate water trapped
				leftIndex := stack[len(stack)-1]
				minHeight := min(height[leftIndex], height[rightIndex])
				width := rightIndex - leftIndex - 1
				totalWaterTrapped += width * (minHeight - height[topIndex])
			}
		}
		stack = append(stack, rightIndex)
	}

	return totalWaterTrapped
}
