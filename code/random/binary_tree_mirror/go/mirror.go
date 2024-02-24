package main

type Node struct {
	val   int
	right *Node
	left  *Node
}

func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func IsMirror(root *Node) bool {
	var recursive func(left *Node, right *Node) bool

	recursive = func(left *Node, right *Node) bool {

		if (left.isLeaf() && !right.isLeaf()) || (!left.isLeaf() && right.isLeaf()) {
			return false
		}

		if left.isLeaf() && right.isLeaf() {
			return left.val == right.val
		}

		fromLeft := recursive(left.left, right.right)
		fromRight := recursive(left.right, right.left)

		return fromLeft && fromRight
	}

	// from root
	return recursive(root.left, root.right)
}
