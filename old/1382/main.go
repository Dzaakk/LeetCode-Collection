func balanceBST(root *TreeNode) *TreeNode {
	sorted := getSorted(root)
	return buildBST(sorted)
}
func buildBST(nodes []*TreeNode) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	mid := nodes[len(nodes)/2]
	mid.Left = buildBST(nodes[:len(nodes)/2])
	mid.Right = buildBST(nodes[(len(nodes)/2)+1:])

	return mid
}

func getSorted(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	sortedLeft := getSorted(root.Left)
	sortedLeft = append(sortedLeft, root)
	sortedRight := getSorted(root.Right)
	return append(sortedLeft, sortedRight...)
}