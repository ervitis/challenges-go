package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	n := 0
	count(root, &n)
	return n
}

func count(root *TreeNode, n *int) {
	if root == nil {
		return
	}
	*n++
	count(root.Left, n)
	count(root.Right, n)
}

func main() {
	fmt.Println(countNodes(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}))
}
