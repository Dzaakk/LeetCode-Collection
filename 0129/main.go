func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}
func dfs(node *TreeNode, number int) int {
	if node == nil {
		return 0
	}

	number = number*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return number
	}

	return dfs(node.Left, number) + dfs(node.Right, number)
}