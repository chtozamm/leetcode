// 238. Product of Array Except Self
package product_of_array

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// Calculate the products of elements to the left of each index
	leftProduct := 1
	for i := 0; i < n; i++ {
		res[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Calculate the products of elements to the right of each index
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return res
}
