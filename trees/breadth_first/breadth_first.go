package trees

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int
}

func bfs(head *BinaryNode, needle int) bool {
	q := []*BinaryNode{head}
	for q != nil {
		curr := q[0]
		q = q[1:]
		if curr == nil {
			break
		}
		if curr.value == needle {
			return true
		}
		q = append(q, curr.left)
		q = append(q, curr.right)
	}
	return false
}
