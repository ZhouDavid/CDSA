package linkedlist

import (
	"errors"
)

type ListNode struct {
	Value interface{}
	prev  *ListNode
	next  *ListNode
}
type List struct {
	size int
	head *ListNode
	tail *ListNode
	cur  *ListNode
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		Value: v,
		prev:  nil,
		next:  nil,
	}
}
func (node *ListNode) Next() *ListNode {
	return node.next
}

func (node *ListNode) Prev() *ListNode {
	return node.prev
}

func (node *ListNode) InsertAfter(value interface{}) *ListNode {
	newNode := NewListNode(value)
	succ := node.next
	newNode.prev = node
	newNode.next = succ
	node.next = newNode
	if succ != nil {
		succ.prev = newNode
	}
	return newNode
}

func (node *ListNode) InsertBefore(value interface{}) *ListNode {
	prev := node.prev
	newNode := NewListNode(value)
	newNode.next = node
	newNode.prev = prev
	node.prev = newNode
	if prev != nil {
		prev.next = newNode
	}
	return newNode
}

func New() *List {
	head := NewListNode(nil)
	tail := NewListNode(nil)
	head.next = tail
	tail.prev = head
	return &List{
		size: 0,
		head: head,
		tail: tail,
	}
}
func NewList(size int, value interface{}) (*List, error) {
	if size <= 0 {
		return nil, errors.New("size should be positive integer")
	}
	head := NewListNode(nil)
	tail := NewListNode(nil)
	head.next = tail
	tail.prev = head
	l := &List{
		size: 0,
		head: head,
		tail: tail,
	}

	for i := 0; i < size; i++ {
		l.PushBack(value)
	}
	return l, nil

}

func (l *List) Size() int {
	return l.size
}

func (l *List) Empty() bool {
	if l.size == 0 {
		return true
	}
	return false
}

func (l *List) Front() *ListNode {
	return l.head.next
}

func (l *List) Back() *ListNode {
	return l.tail.prev
}

func (l *List) PushFront(value interface{}) *ListNode {
	l.size++
	return l.head.InsertAfter(value)
}

func (l *List) PushBack(value interface{}) *ListNode {
	l.size++
	return l.tail.InsertBefore(value)
}

func (l *List) Insert(index int, value interface{}) (*ListNode, error) {
	if index > l.size || index < 0 {
		return nil, errors.New("illegal index, out of bounds")
	}
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.InsertBefore(value), nil
}
