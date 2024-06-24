func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	leftTail := dfs(node.Left)
	rightTail := dfs(node.Right)

	if node.Left != nil {
		leftTail.Right = node.Right
		node.Right = node.Left
		node.Left = nil
	}

	if rightTail != nil {
		return rightTail
	} else if leftTail != nil {
		return leftTail
	}

	return node
}