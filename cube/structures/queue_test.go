package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()

	ast.True(q.IsEmpty(), "Queue should be empty")

	start, end := 0, 100

	for i := start; i < end; i++ {
		q.Push(i)
		ast.Equal(i-start+1, q.Len(), "Queue length should be %d", i-start+1)
	}

	for i := start; i < end; i++ {
		ast.Equal(i, q.Pop(), "Queue should return %d", i)
		ast.Equal(end-i-1, q.Len(), "Queue length should be %d", end-i-1)
	}

	ast.True(q.IsEmpty(), "Queue should be empty after all items are popped")
}
