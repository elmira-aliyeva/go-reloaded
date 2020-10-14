package student

//BTreeLevelCount lol
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left = 0
	var right = 0
	if root.Right != nil {
		right = BTreeLevelCount(root.Right)
	}
	if root.Left != nil {
		left = BTreeLevelCount(root.Left)
	}
	if right > left {
		return right + 1
	}
	return left + 1

}

//ApplyforLevel lol
func ApplyforLevel(i int, root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 1 {
		f(root.Data)
	} else if i > 1 {
		ApplyforLevel(i-1, root.Left, f)
		ApplyforLevel(i-1, root.Right, f)
	}

}

//BTreeApplyByLevel lol
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	height := BTreeLevelCount(root)
	for i := 1; i <= height; i++ {
		ApplyforLevel(i, root, f)
	}
}
