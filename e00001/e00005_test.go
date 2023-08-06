package e00001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 876.  Given the head of a singly linked list, return the middle of the
// linked list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	middleNode := head
	currentNode := head
	counter := 0
	middleIndex := 0

	for currentNode.Next != nil {
		counter++
		currentNode = currentNode.Next
		if middleIndex*2 < counter {
			middleNode = middleNode.Next
			middleIndex++
		}
	}

	return middleNode
}

func Test_middleNodeOfNil(t *testing.T) {
	var node *ListNode
	assert.Equal(t, node, middleNode(nil))

	head := ListNode{1, nil}
	assert.Equal(t, &head, middleNode(&head))

	secondNode := ListNode{2, nil}
	head.Next = &secondNode
	assert.Equal(t, 2, middleNode(&head).Val)

	thirdNode := ListNode{3, nil}
	secondNode.Next = &thirdNode
	assert.Equal(t, 2, middleNode(&head).Val)

	prevNode := secondNode.Next
	{
		nextNode := ListNode{4, nil}
		prevNode.Next = &nextNode
		prevNode = prevNode.Next
		assert.Equal(t, 3, middleNode(&head).Val)
	}

	{
		nextNode := ListNode{5, nil}
		prevNode.Next = &nextNode
		prevNode = prevNode.Next
		assert.Equal(t, 3, middleNode(&head).Val)
	}
	{
		nextNode := ListNode{6, nil}
		prevNode.Next = &nextNode
		prevNode = prevNode.Next
		assert.Equal(t, 4, middleNode(&head).Val)
	}
}
