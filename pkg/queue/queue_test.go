package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue[int]()

	assert.NotNil(t, queue)
	assert.Empty(t, queue.GetItems(), "")
}

func TestEnqueue(t *testing.T) {
	queue := NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)

	assert.NotNil(t, queue)

	items := queue.GetItems()

	assert.NotEmpty(t, items)
	assert.Contains(t, items, 1)
	assert.Contains(t, items, 2)
}

func TestPeekWithAQueueWithItems(t *testing.T) {
	queue := NewQueue[int]()
	assert.NotNil(t, queue)

	queue.Enqueue(1)
	item, hasItem := queue.Peek()

	assert.Equal(t, 1, item)
	assert.Equal(t, true, hasItem)
}

func TestPeekWithAEmptyQueue(t *testing.T) {
	queue := NewQueue[int]()
	assert.NotNil(t, queue)

	item, hasItem := queue.Peek()

	assert.Equal(t, 0, item)
	assert.Equal(t, false, hasItem)
}

func TestDequeueInQueueWithItems(t *testing.T) {
	queue := NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)

	assert.NotNil(t, queue)

	item, hasItems := queue.Dequeue()
	assert.Equal(t, 1, item)
	assert.Equal(t, true, hasItems)

	item, hasItems = queue.Dequeue()

	assert.Equal(t, 2, item)
	assert.Equal(t, true, hasItems)

	item, hasItems = queue.Dequeue()

	assert.Equal(t, 0, item)
	assert.Equal(t, false, hasItems)
	assert.Empty(t, queue.last)
}

func TestDequeueWithAEmptyQueue(t *testing.T) {
	queue := NewQueue[int]()

	assert.NotNil(t, queue)

	item, hasItems := queue.Dequeue()
	assert.Equal(t, 0, item)
	assert.Equal(t, false, hasItems)
}

func TestIsEmptyWithAEmptyQueue(t *testing.T) {
	queue := NewQueue[int]()

	assert.NotNil(t, queue)
	assert.True(t, queue.IsEmpty())
}

func TestIsEmptyInAQueueWithItems(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)

	assert.NotNil(t, queue)
	assert.False(t, queue.IsEmpty())
}

func TestGetItemsInAQueueWithItems(t *testing.T) {
	queue := NewQueue[int]()

	queue.Enqueue(0)
	queue.Enqueue(1)

	items := queue.GetItems()

	assert.NotNil(t, queue)
	assert.NotEmpty(t, items)
	assert.Contains(t, items, 0)
	assert.Contains(t, items, 1)
}

func TestGetItemsWithAEmptyQueue(t *testing.T) {
	queue := NewQueue[int]()

	items := queue.GetItems()

	assert.NotNil(t, queue)
	assert.Empty(t, items)
}
