package structures

type Queue struct {
	Items []interface{}
}

// Create a new queue that is empty
func NewQueue() *Queue {
	return &Queue{Items: make([]interface{}, 0)}
}

// Add a new item to the end of the queue
func (q *Queue) Push(item interface{}) {
	q.Items = append(q.Items, item)
}

// Remove and return the front item from the queue
func (q *Queue) Pop() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}

// Return the number of items in  the queue
func (q *Queue) Len() int {
	return len(q.Items)
}

// Test to see whether the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

// Return the front item from the queue without removing it
func (q *Queue) Front() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	return q.Items[0]
}

// Return the last item from the queue without removing it
func (q *Queue) Back() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	return q.Items[len(q.Items)-1]
}
