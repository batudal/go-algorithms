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
	path = append(path, curr.value)
	walk(curr.left, path)
	walk(curr.right, path)
	return path
}

func pre_order_search(head *BinaryNode) []int {
	return walk(head, []int{})
}
