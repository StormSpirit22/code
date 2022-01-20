package main

import "fmt"

func main() {
	//root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	//flatten(root)

	//buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3})
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(minDepth(root))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	left := helper(root.Left)
	right := helper(root.Right)
	if left != nil {
		root.Right = left
		tmp := root.Right
		fmt.Println(tmp.Val)
		for tmp != nil && tmp.Right != nil {
			tmp = tmp.Right
		}
		if tmp != nil {
			tmp.Right = right
		}
	} else {
		root.Right = right
	}
	root.Left = nil
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := root.Right
	if root.Left != nil {
		root.Left = helper(root.Left)
		root.Right = root.Left
		root.Left = nil
		root.Right.Right = right
	}
	return root

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	n := len(postorder)
	root := &TreeNode{Val: postorder[n-1]}
	var index int
	for i := range inorder {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	postorder = postorder[:n-1]
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:])

	return root
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		count++
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return count
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}
	return count
}