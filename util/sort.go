package util

import (
	"GoSTL/heap"
	"sync"
)

type Sortable interface {
	Len() int           // Len is the number of elements in the collection
	Less(i, j int) bool // Whether the element with index i should sort before the element with index j.
	Swap(i, j int)      // Swap swaps the elements with indexes i and j.
}

// Quick Sort
func QuickSort(data Sortable) {
	quickSort(data, 0, data.Len())
}
func quickSort(data Sortable, lo int, hi int) {
	if hi-lo <= 1 {
		return
	}
	i := lo
	j := hi - 1
	for i < j {
		for i < j && !data.Less(j, lo) {
			j--
		}
		for i < j && !data.Less(lo, i) {
			i++
		}
		if i < j {
			data.Swap(i, j)
		}
	}
	data.Swap(i, lo)
	quickSort(data, lo, i)
	quickSort(data, i+1, hi)
}

func HeapSort(data heap.Heapable) {
	h := heap.NewMaxHeap(data)
	n := data.Len() - 1
	for i := n; i >= 1; i-- {
		h.Swap(0, i)
		h.SiftDown(0, i)
	}
}

// multi threads
func MultiThreadQuickSort(data Sortable) {
	multiThreadQuickSort(data, 0, data.Len())
}
func multiThreadQuickSort(data Sortable, lo int, hi int) {
	if hi-lo < 1000 {
		quickSort(data, lo, hi)
		return
	}
	i := lo
	j := hi - 1
	for i < j {
		for i < j && !data.Less(j, lo) {
			j--
		}
		for i < j && !data.Less(lo, i) {
			i++
		}
		if i < j {
			data.Swap(i, j)
		}
	}
	data.Swap(i, lo)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		multiThreadQuickSort(data, lo, i)
	}()
	multiThreadQuickSort(data, i+1, hi)
	wg.Wait()
}

// Insertion Sort
func insertionSort(data Sortable, lo int, hi int) {
	for i := lo + 1; i < hi; i++ {
		for j := i; j > lo && data.Less(j, j-1); j-- {
			data.Swap(j-1, j)
		}
	}
}

func InsertionSort(data Sortable) {
	insertionSort(data, 0, data.Len())
}
