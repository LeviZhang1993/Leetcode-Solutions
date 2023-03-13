package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func match(l, r *TreeNode) bool {
	if l == nil {
		return r == nil
	}
	if r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return match(l.Left, r.Right) && match(l.Right, r.Left)
}

func isSymmetric(root *TreeNode) bool {
	return match(root.Left, root.Right)
}

func main() {
}
