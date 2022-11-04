package main

import "math/rand"

//https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rval := preorder[0]
	i := 0
	for idx, val := range inorder {
		if rval == val {
			i = idx
			break
		}
	}
	left := buildTree(preorder[1:i+1], inorder[:i])
	right := buildTree(preorder[i+1:], inorder[i+1:])
	node := &TreeNode{
		Val:   rval,
		Left:  left,
		Right: right,
	}
	rand.Int()
	return node
}
