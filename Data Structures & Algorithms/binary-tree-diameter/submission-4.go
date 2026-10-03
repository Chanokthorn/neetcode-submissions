/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  f(n) -> maxCurrActiveLen, maxLen

func diameterOfBinaryTree(root *TreeNode) int {
    _, res := f(root, false)
	return res
}

func f(n *TreeNode, isChild bool) (int, int) {
	if n == nil {
		return 0, 0
	}
	lMaxActive, lMax := f(n.Left, true)
	rMaxActive, rMax := f(n.Right, true)
	maxChildActive := int(math.Max(float64(lMaxActive), float64(rMaxActive)))
	maxChild := int(math.Max(float64(lMax), float64(rMax)))

	var maxCurrActive int

	if isChild {
		maxCurrActive = 1 + maxChildActive
	} else {
		maxCurrActive = maxChildActive
	}

	inactiveCandidate := lMaxActive + rMaxActive
	if inactiveCandidate > maxChild {
		maxChild = inactiveCandidate
	}

	if maxCurrActive > maxChild  {
		return maxCurrActive, maxCurrActive
	}

	return maxCurrActive, maxChild
}
