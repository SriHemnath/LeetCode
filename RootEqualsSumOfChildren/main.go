package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(checkTree(&TreeNode{
		5,
		&TreeNode{
			3,
			nil,
			nil,
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}))
}

func checkTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sum := 0
	if root.Left != nil {
		sum += root.Left.Val
	}
	if root.Right != nil {
		sum += root.Right.Val
	}
	return sum == root.Val
}
