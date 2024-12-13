package utils

import (
	"math"
	"sort"
)

// WithinTolerance checks if two floating-point numbers are within a specified tolerance.
// It returns true if the absolute difference between the numbers is less than the tolerance,
// or if the relative difference is within the tolerance for non-zero values of b.
func WithinTolerance(a, b, e float64) bool {
	// Handle exact equality
	if a == b {
		return true
	}

	// Calculate the absolute difference
	d := math.Abs(a - b)

	// Handle the case where b is zero
	if b == 0 {
		return d < e
	}

	// Compare the relative difference
	return d/math.Abs(b) < e
}

// SlicesEqualUnordered checks if two slices are equal ignoring the order.
func SlicesEqualUnordered(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	counts := make(map[int]int)

	for _, num := range slice1 {
		counts[num]++
	}

	for _, num := range slice2 {
		counts[num]--
		if counts[num] < 0 {
			return false
		}
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func ReverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// SplitIntReverse splits integer into a slice of single digits.
// Note that the order of digits is reversed.
func SplitIntReverse(n int) []int {
	s := []int{}
	for n > 0 {
		s = append(s, n%10)
		n = n / 10
	}
	return s
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func AreLinkedListsEqual(l1 *ListNode, l2 *ListNode) bool {
	// Traverse both linked lists simultaneously
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	// Check if both linked lists have reached the end
	return l1 == nil && l2 == nil
}

func SliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, num := range nums[1:] {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return head
}

func SlicesEqualUnordered2D(slice1, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	// Sort inner slices
	for i := range slice1 {
		sort.Ints(slice1[i])
	}
	for i := range slice2 {
		sort.Ints(slice2[i])
	}

	// Sort the outer slices based on the sorted inner slices
	sort.Slice(slice1, func(i, j int) bool {
		return Less(slice1[i], slice1[j])
	})
	sort.Slice(slice2, func(i, j int) bool {
		return Less(slice2[i], slice2[j])
	})

	// Compare sorted slices
	for i := range slice1 {
		for j := range slice1[i] {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}
	}

	return true
}

func Less(slice1, slice2 []int) bool {
	for i := 0; i < len(slice1) && i < len(slice2); i++ {
		if slice1[i] < slice2[i] {
			return true
		}
		if slice1[i] > slice2[i] {
			return false
		}
	}
	return len(slice1) < len(slice2)
}

func LinkedListToSlice(l *ListNode) []int {
	nums := make([]int, 0)
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}
