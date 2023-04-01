package dfs_on_bst

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int
}

func Dfs(head *BinaryNode, needle int) bool {
	return search(head, needle)
}

func search(curr *BinaryNode, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.value == needle {
		return true
	}

	if curr.value > needle {
		return search(curr.right, needle)
	}
	return search(curr.left, needle)
}
