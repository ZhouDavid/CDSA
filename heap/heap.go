package heap

import (
	"GoSTL/util"
	"reflect"
)

type Heap struct {
	elems []interface{}
}

func New() *Heap {
	return &Heap{
		elems: make([]interface{}, 0, 1024), // default capacity is 1024
	}
}

func NewHeapFromSlice(slice interface{}) *Heap {
	s := reflect.ValueOf(slice)
	var elems []interface{}
	if s.Kind() == reflect.Slice {
		elems = make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			elems[i] = s.Index(i).Interface()
		}
	}
	n := s.Len()
	h := &Heap{
		elems: elems,
	}
	for i := n/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *Heap) swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *Heap) Len() int {
	return len(h.elems)
}

// siftdown the element at index i,
// return whether it's going down
// only support minimum heap
func (h *Heap) siftDown(i int) bool {
	n := h.Len()
	i0 := i
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // j is the smaller value of i's two children
		if j2 := j1 + 1; j2 < n && util.Less(&h.elems[j2], &h.elems[j1]) {
			j = j2 // right child is smaller
		}
		if util.Less(&h.elems[j], &h.elems[i]) { // smaller child is smaller than i
			h.swap(i, j)
			i = j
		} else {
			break // i is the smallest, break
		}
	}
	return i > i0
}

// siftup the element at index i
// return whether it's up
func (h *Heap) siftUp(i int) bool {
	n := h.Len()
	i0 := i
	j := (i+1)/2 - 1
	for {
		if j < 0 || j >= n {
			break
		}
		if util.Less(&h.elems[i], &h.elems[j]) {
			h.swap(i, j)
		} else {
			break
		}
	}
	return i < i0
}

func (h *Heap) Pop() interface{} {
	n := h.Len()
	e := h.elems[0]
	h.elems[0] = h.elems[n-1]
	h.elems = h.elems[:n-1]
	h.siftDown(0)
	return e
}

func (h *Heap) Push(e interface{}) {
	h.elems = append(h.elems, e)
	n := h.Len()
	h.siftUp(n - 1)
}
