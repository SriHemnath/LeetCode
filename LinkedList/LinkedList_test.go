package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_AddValue(t *testing.T) {
	list := NewList[int]()

	result, err := list.GetValueAt(1)
	assert.ErrorContains(t, err, "Incorrect position. Total nodes in element is ")
	assert.Equal(t, 0, result)
}
