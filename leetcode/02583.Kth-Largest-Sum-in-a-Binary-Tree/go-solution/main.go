package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Arr []int64

func (a Arr) Len() int { return len(a) }

func (a Arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Arr) Less(i, j int) bool { return a[i] < a[j] }

func dfs(r *TreeNode, lvl int, x map[int]int64) {
	if r == nil {
		return
	}
	x[lvl] += int64(r.Val)
	dfs(r.Left, lvl+1, x)
	dfs(r.Right, lvl+1, x)
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	x := map[int]int64{}

	dfs(root, 0, x)
	arr := Arr{}
	for _, v := range x {
		arr = append(arr, v)
	}
	sort.Sort(arr)
	if len(arr) < k {
		return -1
	}
	return arr[len(arr)-k]
}

func main() {
}
