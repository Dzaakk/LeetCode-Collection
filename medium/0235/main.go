func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}