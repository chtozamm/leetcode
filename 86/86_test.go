package partition_list

import "testing"

func TestPartition(t *testing.T) {
	testCases := []struct {
		head     *ListNode
		x        int
		expected *ListNode
	}{
		{
			head:     sliceToLinkedList([]int{1, 4, 3, 2, 5, 2}),
			x:        3,
			expected: sliceToLinkedList([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			head:     sliceToLinkedList([]int{2, 1}),
			x:        2,
			expected: sliceToLinkedList([]int{1, 2}),
		},
	}

	for i, tc := range testCases {
		got := partition(tc.head, tc.x)
		if !areLinkedListsEqual(got, tc.expected) {
			t.Errorf("Case %d: got %v, expected %v", i+1, linkedListToSlice(got), linkedListToSlice(tc.expected))
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

func linkedListToSlice(l *ListNode) []int {
	nums := make([]int, 0)
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}
