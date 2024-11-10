package gollections

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPush(t *testing.T) {
	type testCase struct {
		inputs        []int
		expectedSlice []int
	}

	cases := map[string]testCase{
		"push to an empty stack": {
			inputs:        []int{1},
			expectedSlice: []int{1},
		},
		"push multiple values": {
			inputs:        []int{1, 2, 3},
			expectedSlice: []int{1, 2, 3},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, i := range tc.inputs {
				s.Push(i)
			}
			assert.Equal(t, tc.expectedSlice, s.data)
		})
	}
}

func TestStackPop(t *testing.T) {
	type testCase struct {
		inputs        []int
		popValue      int
		expectedErr   error
		expectedSlice []int
	}

	cases := map[string]testCase{
		"pop from a non-empty stack": {
			inputs:        []int{1, 2, 3},
			popValue:      3,
			expectedErr:   nil,
			expectedSlice: []int{1, 2},
		},
		"pop from an empty stack": {
			inputs:      []int{},
			popValue:    0,
			expectedErr: errors.New("empty stack"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, i := range tc.inputs {
				s.Push(i)
			}
			val, err := s.Pop()
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			assert.Equal(t, tc.popValue, val)
			assert.Equal(t, tc.expectedSlice, s.data)
		})
	}
}

func TestStackPeek(t *testing.T) {
	type testCase struct {
		inputs      []int
		peekValue   int
		expectedErr error
	}

	cases := map[string]testCase{
		"peek at a non-empty stack": {
			inputs:      []int{1, 2, 3},
			peekValue:   3,
			expectedErr: nil,
		},
		"peek at an empty stack": {
			inputs:      []int{},
			peekValue:   0,
			expectedErr: errors.New("empty stack"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, i := range tc.inputs {
				s.Push(i)
			}
			val, err := s.Peek()
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			assert.Equal(t, tc.peekValue, val)
		})
	}
}
