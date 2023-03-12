package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue []*ListNode

func (q Queue) Len() int { return len(q) }

func (q Queue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q Queue) Less(i, j int) bool { return q[i].Val < q[j].Val }

func (q *Queue) Push(x interface{}) { *q = append(*q, x.(*ListNode)) }

func (q *Queue) Pop() interface{} {
	ans := (*q)[len(*q)-1]
	*q = (*q)[0 : len(*q)-1]
	return ans
}

func mergeKLists(lists []*ListNode) *ListNode {
	q := &Queue{}
	for _, node := range lists {
		if node != nil {
			heap.Push(q, node)
		}
	}
	h := &ListNode{-1, nil}
	curr := h
	for len(*q) > 0 {
		c := heap.Pop(q).(*ListNode)
		curr.Next = c
		if c != nil && c.Next != nil {
			heap.Push(q, c.Next)
		}
		c.Next = nil
		curr = curr.Next
	}
	return h.Next
}

func main() {
}
