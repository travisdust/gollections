package gollections

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"enqueue one": {
			input: []int{1},
			want:  []int{1},
		},
		"enqueue many": {
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			q := &Queue[int]{}

			for _, inputElement := range tt.input {
				q.Enqueue(inputElement)
			}

			for _, expectedElement := range tt.want {
				value, err := q.Dequeue()
				if err != nil || value != expectedElement {
					t.Errorf("Enqueue() = %v, want %v", value, expectedElement)
				}
			}
		})
	}
}

func TestQueueDequeue(t *testing.T) {
	tests := map[string]struct {
		setup         func(*Queue[int])
		want          int
		expectedError error
	}{
		"dequeue from empty": {
			setup:         func(q *Queue[int]) {},
			want:          0,
			expectedError: errors.New("empty queue"),
		},
		"dequeue from single": {
			setup: func(q *Queue[int]) {
				q.Enqueue(5)
			},
			want:          5,
			expectedError: nil,
		},
		"dequeue from multiple": {
			setup: func(q *Queue[int]) {
				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)
			},
			want:          1,
			expectedError: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			q := &Queue[int]{}
			tt.setup(q)
			got, err := q.Dequeue()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, err, tt.expectedError)
		})
	}
}
