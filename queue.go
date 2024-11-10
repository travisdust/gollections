package gollections

import (
	"errors"
	"sync"
)

type Queue[T any] struct {
	mu   sync.Mutex
	data []T
}

func (q *Queue[T]) Enqueue(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, v)
}

func (q *Queue[T]) Dequeue() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	var v T
	if len(q.data) == 0 {
		return v, errors.New("empty queue")
	}

	v = q.data[0]
	q.data = q.data[1:]
	return v, nil
}

func (q *Queue[T]) Peek() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	var v T
	if len(q.data) == 0 {
		return v, errors.New("empty queue")
	}
	return q.data[0], nil
}
