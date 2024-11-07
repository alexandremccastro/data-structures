package queue

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	first *Node[T]
	last  *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	var node Node[T] = Node[T]{
		value: v,
	}

	if q.first == nil {
		q.first = &node
		q.last = &node

		return
	}

	q.last.next = &node
	q.last = &node
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.first == nil {
		var zeroT T
		return zeroT, false
	}

	return q.first.value, true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.first == nil {
		var zeroT T
		return zeroT, false
	}

	value := q.first.value
	q.first = q.first.next

	if q.first == nil {
		q.last = nil
	}

	return value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue[T]) GetItems() []T {
	var items []T
	current := q.first

	for current != nil {
		items = append(items, current.value)
		current = current.next
	}

	return items
}
