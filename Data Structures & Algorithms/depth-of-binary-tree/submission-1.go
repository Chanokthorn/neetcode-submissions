/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	rDepth := maxDepth(root.Right)
	lDepth := maxDepth(root.Left)
	if rDepth > lDepth {
		return 1 + rDepth
	}
	
	return 1 + lDepth
}
