package trees

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int
}

func walk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}
	walk(curr.left, path)
	walk(curr.right, path)
	path = append(path, curr.value)
	return path
}

func pre_order_search(head *BinaryNode) []int {
	return walk(head, []int{})
}
