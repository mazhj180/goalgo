package trees

type AvlTreeNode struct {
	value  int
	height int
	left   *AvlTreeNode
	right  *AvlTreeNode
}

// 跟新节点高度
func (at *AvlTreeNode) updateHeight() {
	at.height = 1 + max(at.left.height, at.right.height)
}

// 左旋
func (at *AvlTreeNode) leftRotate() (node *AvlTreeNode) {
	r := at.right
	at.right = r.left
	r.left = at

	r.updateHeight()
	at.updateHeight()
	return r
}

// 右旋
func (at *AvlTreeNode) rightRotate() *AvlTreeNode {
	l := at.left
	at.left = l.right
	l.right = at

	l.updateHeight()
	at.updateHeight()
	return l
}

// InsertNode 向AVL树中插入元素
func InsertNode(node *AvlTreeNode, val int) *AvlTreeNode {
	if node == nil {
		return &AvlTreeNode{value: val}
	}
	if val < node.value {
		node.left = InsertNode(node.left, val)
	} else if val > node.value {
		node.right = InsertNode(node.right, val)
	} else {
		return node
	}
	node.updateHeight()
	balance := node.left.height - node.left.height

	//right rotate
	if balance > 1 && val < node.left.value {
		return node.rightRotate()
	}

	//left rotate
	if balance < -1 && val > node.right.value {
		return node.leftRotate()
	}

	//left right rotate
	if balance > 1 && val > node.left.value {
		node.left = node.left.leftRotate()
		return node.rightRotate()
	}

	//right left rotate
	if balance < -1 && val < node.right.value {
		node.right = node.right.rightRotate()
		return node.leftRotate()
	}

	return node
}
