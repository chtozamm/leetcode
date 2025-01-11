package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func listFromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	curr := head

	for _, val := range arr[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return head
}

func arrayFromList(list *ListNode) []int {
	arr := make([]int, 0)
	for list != nil {
		arr = append(arr, list.Val)
		list = list.Next
	}
	return arr
}
