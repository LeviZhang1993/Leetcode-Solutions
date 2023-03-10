package main

import "math/rand"

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
type Solution struct {
	curr *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	ans := this.curr
	h := this.curr
	round := 1
	for h != nil && h.Next != nil {
		h = h.Next
		round += 1
		// get randint [1, round]
		if (rand.Int()%round + 1) <= 1 {
			ans = h
		}
	}
	return ans.Val
}

func main() {

}
