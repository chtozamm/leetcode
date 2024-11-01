// 2095. Delete the Middle Node of a Linked List
package delete_middle

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	// Edge case: if the list is empty or has only one node
	if head == nil || head.Next == nil {
		return nil
	}

	sentinel := &ListNode{Next: head}
	slow := sentinel
	fast := head

	// Move slow one step and fast two steps
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Remove the middle node
	slow.Next = slow.Next.Next

	return sentinel.Next
}
