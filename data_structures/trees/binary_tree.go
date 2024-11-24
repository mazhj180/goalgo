package trees

type binaryTreeNode struct {
	val   any
	left  *binaryTreeNode
	right *binaryTreeNode
}

type BinaryTree struct {
	root *binaryTreeNode
}

// BuildTreeByFloor  层序遍历构建二叉树
func BuildTreeByFloor(values []any) (bt *BinaryTree) {
	if values == nil {
		return nil
	}
	root := &binaryTreeNode{val: values[0]}
	queue := []*binaryTreeNode{root}
	for i := 1; len(queue) > 0 && i < len(values); {
		node := queue[0]
		queue = queue[1:]
		if i < len(values) {
			if values[i] != "#" {
				node.left = &binaryTreeNode{val: values[i]}
				queue = append(queue, node.left)
			}
			i++
		}
		if i < len(values) {
			if values[i] != "#" {
				node.right = &binaryTreeNode{val: values[i]}
				queue = append(queue, node.right)
			}
			i++
		}
	}
	return &BinaryTree{root: root}
}

// FloorTraverse 二叉树的层序遍历(广度优先)
func (bt *BinaryTree) FloorTraverse() (res []any) {
	if bt.root == nil {
		return nil
	}
	queue := []*binaryTreeNode{bt.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return
}

// PreTraverse 前序遍历
func (bt *BinaryTree) PreTraverse() (res []any) {
	if bt.root == nil {
		return res
	}
	var stack []*binaryTreeNode
	node := bt.root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.val)
			stack = append(stack, node)
			node = node.left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.right
	}
	return
}

// MidTraverse 中序遍历
func (bt *BinaryTree) MidTraverse() (res []any) {
	if bt.root == nil {
		return res
	}
	var stack []*binaryTreeNode
	node := bt.root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.left
		}
		node = stack[len(stack)-1]
		res = append(res, node.val)
		stack = stack[:len(stack)-1]
		node = node.right
	}
	return
}

// PostTraverse 后序遍历
func (bt *BinaryTree) PostTraverse() (res []any) {
	if bt.root == nil {
		return res
	}
	var stack []*binaryTreeNode
	node := bt.root
	var lastVisited *binaryTreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.left
		}
		peekNode := stack[len(stack)-1]
		if peekNode.right == nil || lastVisited == peekNode.right {
			res = append(res, peekNode.val)
			lastVisited = peekNode
			stack = stack[:len(stack)-1]
		} else {
			node = peekNode.right
		}
	}
	return
}
