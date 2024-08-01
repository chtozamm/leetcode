package two_nums

import (
	"testing"
)

type testCase struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []testCase{
		{
			l1:       sliceToLinkedList([]int{2, 4, 3}),
			l2:       sliceToLinkedList([]int{5, 6, 4}),
			expected: sliceToLinkedList([]int{7, 0, 8}),
		},
		{
			l1:       sliceToLinkedList([]int{0}),
			l2:       sliceToLinkedList([]int{1}),
			expected: sliceToLinkedList([]int{1}),
		},
		{
			l1:       sliceToLinkedList([]int{0}),
			l2:       sliceToLinkedList([]int{0}),
			expected: sliceToLinkedList([]int{0}),
		},
		{
			l1:       sliceToLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       sliceToLinkedList([]int{9, 9, 9, 9}),
			expected: sliceToLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			l1:       sliceToLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			l2:       sliceToLinkedList([]int{5, 6, 4}),
			expected: sliceToLinkedList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}
	for _, tc := range testCases {
		got := addTwoNumbers(tc.l1, tc.l2)
		if !areLinkedListsEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}

func sliceToLinkedList(nums []int) *ListNode {
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

func areLinkedListsEqual(l1 *ListNode, l2 *ListNode) bool {
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
