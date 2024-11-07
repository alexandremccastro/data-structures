package deque

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type Deque[T any] struct {
	first *Node[T]
	last  *Node[T]
	size  int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) AddFirst(v T) {
	node := &Node[T]{
		value: v,
	}

	if d.first == nil {
		d.first = node
		d.last = node
	} else {
		node.next = d.first
		d.first.prev = node
		d.first = node
	}

	d.size++
}

func (d *Deque[T]) GetFirst() (T, bool) {
	if d.first == nil {
		var zeroT T
		return zeroT, false
	}

	return d.first.value, true
}

func (d *Deque[T]) RemoveFirst() {
	if d.first == nil {
		return
	}

	if d.first == d.last {
		d.first = nil
		d.last = nil
	} else {
		d.first = d.first.next
		d.first.prev = nil
	}

	d.size--
}

func (d *Deque[T]) AddLast(v T) {
	node := &Node[T]{
		value: v,
	}

	if d.first == nil {
		d.first = node
		d.last = node
	} else {
		node.prev = d.last
		d.last.next = node
		d.last = node
	}

	d.size++
}

func (d *Deque[T]) GetLast() (T, bool) {
	if d.last == nil {
		var zeroT T
		return zeroT, false
	}

	return d.last.value, true
}

func (d *Deque[T]) RemoveLast() {
	if d.last == nil {
		return
	}

	if d.first == d.last {
		d.first = nil
		d.last = nil
	} else {
		d.last = d.last.prev
		d.last.next = nil
	}

	d.size--
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) IsEmpty() bool {
	return d.first == nil
}

func (d *Deque[T]) GetItems() []T {
	var items []T
	current := d.first

	for current != nil {
		items = append(items, current.value)
		current = current.next
	}

	return items
}
