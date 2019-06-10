package list

import (
	"GoSTL/algorithms"
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

func (l *List) Len() int {
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

func (l *List) End() *ListNode {
	return l.tail
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
	l.size++
	return cur.InsertBefore(value), nil
}

func (l *List) Remove(index int, value interface{}) (*ListNode, error) {
	if index >= l.size || index < 0 {
		return nil, errors.New("illegal index, out of bounds")
	}
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	prev := cur.prev
	next := cur.next
	prev.next = next
	next.prev = prev
	l.size--
	return cur, nil
}

func (l *List) RemoveRange(beginIndex int, endIndex int) error {
	if (beginIndex >= l.Len() || beginIndex < 0) ||
		(endIndex > l.Len() || endIndex < 0) || (beginIndex >= endIndex) {
		return errors.New("illegal index, out of bounds")
	}
	cur := l.head.next
	for i := 0; i < beginIndex; i++ {
		cur = cur.next
	}
	beginNode := cur.prev
	for i := beginIndex; i < endIndex; i++ {
		cur = cur.next
	}
	endNode := cur
	beginNode.next = endNode
	endNode.prev = beginNode
	l.size -= endIndex - beginIndex
	return nil
}

func (l *List) Get(i int) *ListNode {
	if i >= l.Len() || i < 0 {
		return nil
	}
	cur := l.head.next
	for i > 0 {
		cur = cur.next
		i--
	}
	return cur
}

func (l *List) Swap(i, j int) {
	if i == j {
		return
	}
	if i > j {
		i, j = j, i
	}
	a := l.Get(i)
	b := l.Get(j)
	if a.next == b || b.next == a {
		aPrev := a.prev
		bNext := b.next
		aPrev.next = b
		bNext.prev = a
		b.prev = aPrev
		a.next = bNext
		a.prev = b
		b.next = a
	} else {
		aPrev := a.prev
		aNext := a.next
		bPrev := b.prev
		bNext := b.next
		aPrev.next = b
		b.prev = aPrev
		b.next = aNext
		aNext.prev = b
		a.prev = bPrev
		a.next = bNext
		bPrev.next = a
		bNext.prev = a
	}

}

func (l *List) Less(i, j int) bool {
	a := l.Get(i)
	b := l.Get(j)
	if comparable, ok := a.Value.(algorithms.Comparable); ok {
		return comparable.Less(b.Value)
	} else {
		// directly compare between basic types
		switch a.Value.(type) {
		case int:
			return a.Value.(int) < b.Value.(int)
		case int8:
			return a.Value.(int8) < b.Value.(int8)
		case int16:
			return a.Value.(int16) < b.Value.(int16)
		case int32:
			return a.Value.(int32) < b.Value.(int32)
		case int64:
			return a.Value.(int64) < b.Value.(int64)
		case float32:
			return a.Value.(float32) < b.Value.(float32)
		case float64:
			return a.Value.(float64) < b.Value.(float64)
		case string:
			return a.Value.(string) < b.Value.(string)
		default:
			panic("element type is not basic type and didn't implement Less interface")
		}
	}

}
