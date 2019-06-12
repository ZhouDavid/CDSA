package main

import (
	"GoSTL/heap"
	"GoSTL/list"
	"GoSTL/vector"
	"fmt"
	"math/rand"
	"sort"
)

type Test struct {
	s string
}

func (t Test) change() {
	t.s = "2"
}

func (t *Test) print() {
	fmt.Println(t.s)
}

func main() {
	// heap.Interface
	// test vector
	arr1 := vector.New()
	for i := 0; i < 5; i++ {
		arr1.PushBack(rand.Float64())
	}
	sort.Sort(arr1)
	fmt.Printf("%v\n", *arr1)
	// test list
	l1 := list.New()
	for i := 0; i < 10; i++ {
		l1.PushBack(rand.Float64())
	}
	l1.Insert(3, 123123.)
	sort.Sort(l1)

	arr2 := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		arr2[i] = 9 - i
	}
	vector.NewVectorFromSlice(arr2)
	h := heap.NewHeapFromSlice(arr2)
	fmt.Println(h.Len())
	for i := 0; i < 10; i++ {
		fmt.Println(h.Pop())
	}
}
