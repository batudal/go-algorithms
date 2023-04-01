package heap

type MinHeap struct {
	length int
	data   []int
}

type Heap interface {
	insert(value int)
	delete() int
}

func (h *MinHeap) Insert(value int) {
	h.data[h.length] = value
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	if h.length == 0 {
		return -1
	}

	out := h.data[0]
	h.length--

	if h.length == 1 {
		h.data = make([]int, 0)
		return out
	}

	h.data[0] = h.data[h.length]
	h.heapifyDown(0)
	return out
}

func (h *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := h.parent(idx)
	pv := h.data[p]
	v := h.data[idx]

	if pv > v {
		h.data[idx] = pv
		h.data[p] = v
		h.heapifyUp(p)
	}
}

func (h *MinHeap) heapifyDown(idx int) {

	li := h.leftChild(idx)
	ri := h.rightChild(idx)

	if idx >= h.length || li > h.length {
		return
	}

	lv := h.data[li]
	rv := h.data[ri]
	v := h.data[idx]

	if lv > rv && v > rv {
		h.data[idx] = rv
		h.data[ri] = v
		h.heapifyDown(ri)
	} else if rv > lv && v > lv {
		h.data[idx] = lv
		h.data[li] = v
		h.heapifyDown(li)
	}

}
