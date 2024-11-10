package gollections

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	mu   sync.Mutex
	data []T
}

// Push adds an element of generic type T to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, v)
}

// Pop removes and returns the element from the top of the stack. Returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var value T

	if len(s.data) == 0 {
		return value, errors.New("empty stack")
	}

	value = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}

// Peek returns the element at the top of the stack without removing it. Returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var value T
	if len(s.data) == 0 {
		return value, errors.New("empty stack")
	}
	value = s.data[len(s.data)-1]
	return value, nil
}
