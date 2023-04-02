package lru

type LRU[K comparable, V any] struct {
	length         int
	capacity       int
	head           *Node[V]
	tail           *Node[V]
	lookup         map[K]*Node[V]
	reverse_lookup map[*Node[V]]K
}

type Node[V any] struct {
	value V
	next  *Node[V]
	prev  *Node[V]
}

func createNode[V any](value V) *Node[V] {
	return &Node[V]{value: value}
}
func (l *LRU[K, V]) update(key K, value V) {
	node := l.lookup[key]
	if node == nil {
		node = createNode(value)
		l.length++
		l.prepend(node)
		l.trimCache()
		l.lookup[key] = node
		l.reverse_lookup[node] = key
	} else {
		l.detach(node)
		l.prepend(node)
		node.value = value
	}

}

func (l *LRU[K, V]) get(key K) V {
	node := l.lookup[key]
	if node == nil {
		return *new(V)
	}
	l.detach(node)
	l.prepend(node)
	return node.value
}

func (l *LRU[K, V]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = l.head.next
	}

	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (l *LRU[K, V]) prepend(node *Node[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}
func (l *LRU[K, V]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)

	key := l.reverse_lookup[tail]
	delete(l.lookup, key)
	delete(l.reverse_lookup, tail)
	l.length--

}
