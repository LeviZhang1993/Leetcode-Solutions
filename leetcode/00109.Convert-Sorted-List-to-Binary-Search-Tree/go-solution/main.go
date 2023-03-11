package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(slow.Next),
	}
}

func main() {
}
