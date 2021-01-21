package chapter4

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func CreateMinTree(input []int) (root *TreeNode) {
	rootIndex := len(input) / 2
	rootNode := TreeNode{}
	rootNode.Value = input[rootIndex]
	rootNode.Left = CreateMinTree(input[:rootIndex])
	rootNode.Right = CreateMinTree(input[rootIndex+1:])
	return &rootNode
}
