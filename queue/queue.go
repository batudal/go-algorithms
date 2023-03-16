package queue

type Queue[T any] struct {
  length int
  head *Node[T]
  tail *Node[T]
}

type Node[T any] struct {
  next *Node[T]
  value T
}

func (q *Queue[T]) enqueue(v T) {
  q.length++
  if (q.length == 1) {
    q.tail.value = v
    q.head.value = v
    return
  }

  q.tail.next.value = v
  q.tail.value = v
}

func (q *Queue[T]) dequeue() T {
  var t T
  if (q.length == 0) {
    return t
  }
  q.length--

  head := q.head
  q.head = q.head.next
  
  return head.value
}

func (q *Queue[T]) peek() T {
  return q.head.value
}
