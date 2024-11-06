package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	assert.NotNil(t, stack, "stack must be defined")
	assert.True(t, stack.IsEmpty(), "initially stack should must be empty")
}

func TestPush(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)

	assert.NotNil(t, stack, "stack must be defined")
	assert.False(t, stack.IsEmpty(), "stack should not be empty")
	assert.NotEmpty(t, stack.GetItems(), "must return all stack items")
	assert.Equal(t, len(stack.GetItems()), 2, "stack items length must be 2")
}

func TestPeekWithAFullStack(t *testing.T) {
	stack := NewStack[int]()
	assert.NotNil(t, stack, "stack must be defined")

	stack.Push(1)

	value, hasItems := stack.Peek()

	assert.True(t, hasItems, "stack must have items to peek")
	assert.Equal(t, 1, value, "peeked value must be 1")
	assert.Equal(t, len(stack.GetItems()), 1, "stack items length must be 1")
}

func TestPeekWithAEmptyStack(t *testing.T) {
	stack := NewStack[int]()
	assert.NotNil(t, stack, "stack must be defined")

	value, hasItems := stack.Peek()

	assert.False(t, hasItems, "stack must not have items to peek")
	assert.Equal(t, 0, value, "peeked value must be zero")
	assert.Equal(t, len(stack.GetItems()), 0, "stack items length must be 0")
}

func TestPopWithAFullStack(t *testing.T) {
	stack := NewStack[int]()
	assert.NotNil(t, stack, "stack must be defined")

	stack.Push(1)
	stack.Push(2)

	assert.NotEmpty(t, stack.GetItems(), "must return all stack items")
	assert.Equal(t, len(stack.GetItems()), 2, "stack items length must be 2")

	value, hasItems := stack.Pop()

	assert.True(t, hasItems, "stack must have items to pop")
	assert.Equal(t, 2, value, "poped value must be equal to 2")
	assert.Equal(t, len(stack.GetItems()), 1, "stack items length must be 1")
}

func TestPopWithAEmptyStack(t *testing.T) {
	stack := NewStack[int]()
	assert.NotNil(t, stack, "stack must be defined")

	value, hasItems := stack.Pop()

	assert.False(t, hasItems, "stack must not have items to pop")
	assert.Equal(t, 0, value, "poped value must be zero")
	assert.Equal(t, len(stack.GetItems()), 0, "stack items length must be 0")
}

func TestGetItemsWithAFullStack(t *testing.T) {
	stack := NewStack[int]()
	assert.NotNil(t, stack, "stack must be defined")

	stack.Push(0)
	stack.Push(1)
	stack.Push(2)

	values := stack.GetItems()

	assert.Equal(t, 3, len(values), "items length must be 3")
	assert.NotEmpty(t, values, "stack must not be empty")
}

func TestGetItemsWithAEmptyStack(t *testing.T) {
	stack := NewStack[int]()
	assert.NotNil(t, stack, "stack must be defined")

	values := stack.GetItems()

	assert.Equal(t, 0, len(values), "items length must be 0")
	assert.Empty(t, values, "stack must be empty")
}
