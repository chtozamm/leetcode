package remove_duplicates_ii

import (
	"reflect"
	"testing"
)

type testCase struct {
	head     *ListNode
	expected []int
}

func TestDeleteDuplicates(t *testing.T) {
	testCases := []testCase{
		{
			head:     sliceToLinkedList([]int{1, 2, 3, 3, 4, 4, 5}),
			expected: []int{1, 2, 5},
		},
		{
			head:     sliceToLinkedList([]int{1, 1, 1, 2, 3}),
			expected: []int{2, 3},
		},
	}

	for _, tc := range testCases {
		got := linkedListToSlice(deleteDuplicates(tc.head))
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}

func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return head
}

func linkedListToSlice(l *ListNode) []int {
	nums := make([]int, 0)
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}
