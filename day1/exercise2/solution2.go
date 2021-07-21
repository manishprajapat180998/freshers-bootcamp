package main

import "fmt"

type TreeNode struct{
	x string
	left *TreeNode
	right *TreeNode
}
func PostorderTraversal(tree *TreeNode) {
	if tree == nil{
		return
	} else{
		PreorderTraversal(tree.left)
		PreorderTraversal(tree.right)
		fmt.Println(tree.x)
	}
}
func PreorderTraversal(tree *TreeNode) {
	if tree == nil{
		return
	} else{
		fmt.Println(tree.x)
		PreorderTraversal(tree.left)
		PreorderTraversal(tree.right)
	}
}
func main(){
	tree := &TreeNode{"+",nil,nil}
	tree.left = &TreeNode{"a",nil, nil}
	tree.right = &TreeNode{"-",nil,nil}
	tree.right.left = &TreeNode{"b",nil, nil}
	tree.right.right = &TreeNode{"c", nil, nil }
	fmt.Println("PreorderTraversal :")
	PreorderTraversal(tree)
	fmt.Println("PostorderTraversal :")
	PostorderTraversal(tree)
}