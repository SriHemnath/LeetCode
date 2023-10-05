package LinkedList

import "testing"

func TestLinkedList_AddValue(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		n    *LinkedList[T]
		args args[T]
	}
	list := NewList[int]()
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "Test1",
			n:    list,
			args: args[int]{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.AddValue(tt.args.value)
		})
	}
}
