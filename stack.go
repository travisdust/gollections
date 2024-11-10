package gollections

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	mu   sync.Mutex
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, v)
}

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
