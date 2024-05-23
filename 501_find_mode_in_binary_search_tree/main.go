package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var modes []int
	var prev *TreeNode
	var maxCount, currCount int

	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if prev != nil && prev.Val == node.Val {
			currCount++
		} else {
			currCount = 1
		}

		if currCount > maxCount {
			maxCount = currCount
			modes = []int{node.Val}
		} else if currCount == maxCount {
			modes = append(modes, node.Val)
		}
		prev = node
		inorder(node.Right)
	}
	inorder(root)
	return modes
}

func main() {
	fmt.Println(findMode(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}))
	fmt.Println(findMode(&TreeNode{Val: 0}))
	fmt.Println(findMode(&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}))
}
