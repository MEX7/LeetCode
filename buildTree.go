package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	for k := range inorder {
		if inorder[k] == preorder[0] { //中序遍历 root (index=k)
			return &TreeNode{
				Val: preorder[0],
				//Val: inorder[root],
				Left:  BuildTree(preorder[1:k+1], inorder[0:k]),
				Right: BuildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}
