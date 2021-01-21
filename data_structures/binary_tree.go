package data_structures

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func (t *BST) Insert(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Value: val}
		return
	}
	insert(t.Root, val)
}

func insert(parent *TreeNode, val int) {
	if parent == nil {
		parent = &TreeNode{Value: val}
		return
	}
	if parent.Value == val {
		return
	}
	if parent.Value < val {
		if parent.Right == nil {
			parent.Right = &TreeNode{Value: val}
			return
		}
		insert(parent.Right, val)
	} else {
		if parent.Left == nil {
			parent.Left = &TreeNode{Value: val}
			return
		}
		insert(parent.Left, val)
	}
}

func (t *BST) Delete(val int) {
	if t.Root == nil {
		return
	}
	t.Root = deleteNode(t.Root, val)
}

func deleteNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Value < val {
		node.Right = deleteNode(node.Right, val)
		return node
	} else if node.Value > val {
		node.Left = deleteNode(node.Left, val)
		return node
	} else {
		if node.Left == nil && node.Right == nil {
			node = nil
			return node
		}
		if node.Left == nil {
			node = node.Right
			return node
		} else if node.Right == nil {
			node = node.Left
			return node
		} else {
			minNode := findMinParent(node.Right)
			node.Value = minNode.Left.Value
			minNode.Left = nil
			return node
		}
	}
}

func findMinParent(node *TreeNode) *TreeNode {
	if node.Left != nil {
		if node.Left.Left != nil {
			return findMinParent(node.Left)
		}
	}
	return node
}
