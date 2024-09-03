package remove_duplicates

import (
	"reflect"
	"testing"
)

type testCase struct {
	head     *ListNode
	expected *ListNode
}

func TestDeleteDuplicates(t *testing.T) {
	testCases := []testCase{
		{
			head:     sliceToListNode([]int{1, 1, 2}),
			expected: sliceToListNode([]int{1, 2}),
		},
		{
			head:     sliceToListNode([]int{1, 1, 2, 3, 3}),
			expected: sliceToListNode([]int{1, 2, 3}),
		},
	}

	for _, tc := range testCases {
		got := deleteDuplicates(tc.head)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}

func sliceToListNode(nums []int) *ListNode {
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
