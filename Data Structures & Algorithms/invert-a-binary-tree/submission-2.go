/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)

	prevL := root.Left
	root.Left = root.Right
	root.Right = prevL
	return root
}
