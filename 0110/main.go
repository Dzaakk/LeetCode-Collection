func isBalanced(root *TreeNode) bool {
	if getDepth(root) < 0 {
		return false
	}

	return true
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if (node.Left == nil) && (node.Right == nil) {
		return 1
	}

	left := getDepth(node.Left)
	right := getDepth(node.Right)

	if (left < 0) || (right < 0) || (math.Abs(float64(left-right)) > 1) {
		return -1
	}

	if left > right {
		return left + 1
	}

	return right + 1
}
