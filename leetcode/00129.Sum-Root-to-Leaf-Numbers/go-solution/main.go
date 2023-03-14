package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(r *TreeNode, sum *int, curr int) {
	if r == nil {
		return
	}
	if r.Left == nil && r.Right == nil {
		*sum += curr
	}
	if r.Left != nil {
		dfs(r.Left, sum, curr*10+r.Val)
	}
	if r.Right != nil {
		dfs(r.Right, sum, curr*10+r.Val)
	}
}

func sumNumbers(root *TreeNode) int {
	var ans *int = new(int)
	*ans = 0
	dfs(root, ans, 0)
	return *ans
}

func main() {
}
