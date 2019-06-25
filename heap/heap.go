package heap

const (
	MIN_HEAP = iota
	MAX_HEAP
)

type Heapable interface {
	Len() int           // Len is the number of elements in the collection
	Less(i, j int) bool // Whether the element with index i should sort before the element with index j.
	Swap(i, j int)      // Swap swaps the elements with indexes i and j.
	Push(interface{})
	Pop() interface{}
}

type Heap struct {
	data     Heapable
	heapType int
}

func NewMinHeap(data Heapable) *Heap {
	n := data.Len()
	h := &Heap{
		data:     data,
		heapType: MIN_HEAP,
	}
	for i := n/2 - 1; i >= 0; i-- {
		h.minSiftDown(i, n)
	}
	return h
}

func NewMaxHeap(data Heapable) *Heap {
	n := data.Len()
	h := &Heap{
		data:     data,
		heapType: MAX_HEAP,
	}
	for i := n/2 - 1; i >= 0; i-- {
		h.maxSiftDown(i, n)
	}
	return h
}

func (h *Heap) Swap(i, j int) {
	h.data.Swap(i, j)
}

func (h *Heap) Pop() interface{} {
	n := h.data.Len() - 1
	h.data.Swap(0, n)
	h.SiftDown(0, n)
	return h.Pop() // pop back
}

func (h *Heap) Push(e interface{}) {
	h.data.Push(e)             // push back
	h.SiftUp(h.data.Len() - 1) //
}

func (h *Heap) SiftDown(i int, n int) bool {
	switch h.heapType {
	case MIN_HEAP:
		return h.minSiftDown(i, n)
	default: // max heap
		return h.maxSiftDown(i, n)
	}
}

func (h *Heap) SiftUp(i int) bool {
	switch h.heapType {
	case MIN_HEAP:
		return h.minSiftUp(i)
	default: // max heap
		return h.maxSiftUp(i)
	}
}

// siftdown the element at index i,
// return whether it's going down
// only support minimum heap
func (h *Heap) minSiftDown(i int, n int) bool {
	i0 := i
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // j is the smaller value of i's two children
		if j2 := j1 + 1; j2 < n && h.data.Less(j2, j1) {
			j = j2 // right child is smaller
		}
		if h.data.Less(j, i) { // smaller child is smaller than i
			h.data.Swap(i, j)
			i = j
		} else {
			break // i is the smallest, break
		}
	}
	return i > i0
}

// siftup the element at index i
// return whether it's up
func (h *Heap) minSiftUp(i int) bool {
	n := h.data.Len()
	i0 := i
	for {
		j := (i+1)/2 - 1
		if j < 0 || j >= n {
			break
		}
		if h.data.Less(i, j) {
			h.data.Swap(i, j)
			i = j
		} else {
			break
		}
	}
	return i < i0
}

func (h *Heap) maxSiftDown(i int, n int) bool {
	i0 := i
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // j is the larger value of i's two children
		if j2 := j1 + 1; j2 < n && h.data.Less(j1, j2) {
			j = j2 // right child is larger
		}
		if h.data.Less(i, j) { // larger child is largert than i
			h.data.Swap(i, j)
			i = j
		} else {
			break // i is the largest, break
		}
	}
	return i > i0
}

// siftup the element at index i
// return whether it's up
func (h *Heap) maxSiftUp(i int) bool {
	n := h.data.Len()
	i0 := i
	for {
		j := (i+1)/2 - 1
		if j < 0 || j >= n {
			break
		}
		if h.data.Less(j, i) { //parent is smaller than child, need swap
			h.data.Swap(i, j)
			i = j
		} else {
			break // parent is larget, no need swap
		}
	}
	return i < i0
}
