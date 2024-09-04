// 82. Remove Duplicates from Sorted List II
package remove_duplicates_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Sentinel node helps to handle edge case when the head itself is a duplicate
	sentinel := &ListNode{Next: head}
	prev := sentinel
	cur := head

	for cur != nil {
		// Check if the current value is a duplicate
		if cur.Next != nil && cur.Val == cur.Next.Val {
			// Skip all nodes with the same value
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			// Link the previous node to the next distinct value
			prev.Next = cur.Next
		} else {
			// No duplicates, move the previous pointer
			prev = cur
		}
		// Move to the next node
		cur = cur.Next
	}

	return sentinel.Next // Return the new head, which is the next of sentinel
}
