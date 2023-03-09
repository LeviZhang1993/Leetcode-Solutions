package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	q, s := head.Next.Next, head.Next
	for q != nil && q.Next != nil && q != s {
		q = q.Next.Next
		s = s.Next
	}
	if q == nil || q.Next == nil {
		return nil
	}
	curr := head
	for curr != q {
		q = q.Next
		curr = curr.Next
	}
	return curr
}

func main() {
	println(detectCycle(&ListNode{-1, nil}) == nil)
	node1, node2, node3, node4 := &ListNode{1, nil}, &ListNode{2, nil}, &ListNode{3, nil}, &ListNode{4, nil}
	node1.Next, node2.Next, node3.Next = node2, node3, node4
	node4.Next = node2
	println(detectCycle(node1) == node2)
}
