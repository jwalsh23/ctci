package data_structures

type treeNode struct{
	Value int
	left *treeNode
	right *treeNode
}

type BST struct {
	root *treeNode
}

func(t *BST) Insert(val int) {
	if t.root == nil {
		t.root = &treeNode{Value: val}
		return
	}
	insert(t.root, val)
}

func insert(parent *treeNode, val int) {
	if parent == nil {
		parent = &treeNode{Value: val}
		return
	}
	if parent.Value == val {
		return
	}
	if parent.Value < val {
		if parent.right == nil {
			parent.right = &treeNode{Value: val}
			return
		}
		insert(parent.right, val)
	} else {
		if parent.left == nil {
			parent.left = &treeNode{Value: val}
			return
		}
		insert(parent.left, val)
	}
}

func(t *BST) Delete(val int) {
	if t.root == nil {
		return
	}
	t.root = deleteNode(t.root, val)
}

func deleteNode(node *treeNode, val int) *treeNode {
	if node == nil {
		return nil
	}
 	if node.Value < val {
		node.right = deleteNode(node.right, val)
		return node
	} else if node.Value > val{
		node.left = deleteNode(node.left, val)
		return node
	} else {
		if node.left == nil  && node.right == nil {
			node = nil
			return node
		}
		if node.left == nil {
			node = node.right
			return node
		} else if node.right == nil {
			node = node.left
			return node
		} else {
			minNode := findMinParent(node.right)
			node.Value = minNode.left.Value
			minNode.left = nil
			return node
		}
	}
}

func findMinParent(node *treeNode) *treeNode {
	if node.left != nil {
		if node.left.left != nil {
			return findMinParent(node.left)
		}
	}
	return node
}
