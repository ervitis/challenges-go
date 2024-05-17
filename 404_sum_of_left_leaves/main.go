package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var res int
	leftLeaves(root, &res)
	return res
}

func leftLeaves(node *TreeNode, s *int) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		*s += node.Left.Val
	}
	leftLeaves(node.Right, s)
	leftLeaves(node.Left, s)
}

func main() {
	fmt.Println(sumOfLeftLeaves(&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}))
	fmt.Println(sumOfLeftLeaves(&TreeNode{Val: 1}))
}
