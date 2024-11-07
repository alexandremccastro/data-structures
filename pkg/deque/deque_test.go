package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	deque := NewDeque[int]()

	assert.NotNil(t, deque)
	assert.Equal(t, 0, deque.Size())
}

func TestAddFirst(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddFirst(1)
	deque.AddFirst(2)

	items := deque.GetItems()

	assert.NotNil(t, deque)
	assert.Equal(t, 2, items[0])
	assert.Equal(t, 1, items[1])
	assert.Equal(t, 2, deque.Size())
}

func TestGetFirstInANonEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddFirst(1)
	val, hasItems := deque.GetFirst()

	assert.NotNil(t, deque)
	assert.True(t, hasItems)
	assert.Equal(t, 1, val)
	assert.Equal(t, 1, deque.Size())
}

func TestGetFirstInAEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	val, hasItems := deque.GetFirst()

	assert.NotNil(t, deque)
	assert.False(t, hasItems)
	assert.Equal(t, 0, val)
	assert.Equal(t, 0, deque.Size())
}

func TestRemoveFirstInAEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	deque.RemoveFirst()

	assert.NotNil(t, deque)
	assert.Empty(t, deque.GetItems())
}

func TestRemoveFirstInANonEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddFirst(1)
	deque.RemoveFirst()

	assert.NotNil(t, deque)
	assert.Equal(t, 0, deque.Size())

	deque.AddLast(1)
	deque.AddFirst(2)
	deque.RemoveFirst()

	items := deque.GetItems()

	assert.Equal(t, 1, deque.Size())
	assert.Equal(t, 1, items[0])
}

func TestAddLast(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddLast(1)
	deque.AddLast(2)

	items := deque.GetItems()

	assert.NotNil(t, deque)
	assert.Equal(t, 1, items[0])
	assert.Equal(t, 2, items[1])
	assert.Equal(t, 2, deque.Size())
}

func TestGetLastInANonEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddLast(1)
	val, hasItems := deque.GetLast()

	assert.NotNil(t, deque)
	assert.True(t, hasItems)
	assert.Equal(t, 1, val)
	assert.Equal(t, 1, deque.Size())
}

func TestGetLastInAEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	val, hasItems := deque.GetLast()

	assert.NotNil(t, deque)
	assert.False(t, hasItems)
	assert.Equal(t, 0, val)
	assert.Equal(t, 0, deque.Size())
}

func TestIsEmptyInANonEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddFirst(1)

	assert.NotNil(t, deque)
	assert.False(t, deque.IsEmpty())
}

func TestIsEmptyInAEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()
	assert.NotNil(t, deque)
	assert.True(t, deque.IsEmpty())
}

func TestRemoveLastInAEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	deque.RemoveLast()

	assert.NotNil(t, deque)
	assert.Empty(t, deque.GetItems())
}

func TestRemoveLastInANonEmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddFirst(1)
	deque.RemoveLast()

	assert.NotNil(t, deque)
	assert.Equal(t, 0, deque.Size())

	deque.AddLast(1)
	deque.AddFirst(2)
	deque.RemoveFirst()

	items := deque.GetItems()

	assert.Equal(t, 1, deque.Size())
	assert.Equal(t, 1, items[0])
}

func TestGetItems(t *testing.T) {
	deque := NewDeque[int]()

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddLast(3)
	deque.AddLast(4)
	deque.RemoveLast()

	items := deque.GetItems()

	assert.NotNil(t, deque)
	assert.Equal(t, 2, items[0])
	assert.Equal(t, 1, items[1])
	assert.Equal(t, 3, items[2])
	assert.Equal(t, 3, deque.Size())
}
