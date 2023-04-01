package compare_binary

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int
}

func Compare(a *BinaryNode, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return Compare(a.left, b.left) && Compare(a.right, b.right)
}
