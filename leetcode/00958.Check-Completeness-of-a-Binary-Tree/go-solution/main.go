package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	foundNull := false
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr == nil {
			foundNull = true
		} else {
			if foundNull {
				return false
			}
			q = append(q, curr.Left)
			q = append(q, curr.Right)
		}
	}
	return true
}

func main() {
}
