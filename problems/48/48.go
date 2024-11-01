// 48. Rotate Image
package rotate_image

func rotate(matrix [][]int) {
	n := len(matrix)

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - layer - 1
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]

			// left -> top
			matrix[first][i] = matrix[last-offset][first]

			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]

			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]

			// top -> right
			matrix[i][last] = top
		}
	}
}
