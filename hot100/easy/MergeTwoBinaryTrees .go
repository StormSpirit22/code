package easy
//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。 
//
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点
//。 
//
// 示例 1: 
//
// 
//输入: 
//	Tree 1                     Tree 2                  
//          1                         2                             
//         / \                       / \                            
//        3   2                     1   3                        
//       /                           \   \                      
//      5                             4   7                  
//输出: 
//合并后的树:
//	     3
//	    / \
//	   4   5
//	  / \   \ 
//	 5   4   7
// 
//
// 注意: 合并必须从两个树的根节点开始。 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 751 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return traverse(root1, root2)
}

func traverse(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	root := &TreeNode{}
	if root1 != nil && root2 != nil {
		root.Val = root1.Val + root2.Val
		root.Left = traverse(root1.Left, root2.Left)
		root.Right = traverse(root1.Right, root2.Right)
	} else if root1 == nil {
		root = root2
	} else {
		root = root1
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
