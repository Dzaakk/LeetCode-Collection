package main

type Result struct {
	balanced bool
	height   int
}

func isBalanced(root *TreeNode) bool {
	return dfs(root).balanced
}

func dfs(root *TreeNode) Result {
	if root == nil {
		return Result{true, 0}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	balanced := left.balanced && right.balanced && abs(left.height-right.height) <= 1
	return Result{balanced, 1 + max(left.height, right.height)}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
