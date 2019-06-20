package main

import (
	"GoSTL/heap"
	"GoSTL/vector"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	arrLen := 235098309
	v := vector.NewVector(arrLen, 0)
	for i := 0; i < arrLen; i++ {
		v[i] = rand.Float64()
	}
	t1 := time.Now()
	sort.Sort(v)
	t2 := time.Now()
	fmt.Printf("standard sort:%d\n", t2.Sub(t1))
	h := heap.NewHeapFromSlice(v)
	for h.Len() != 0 {
		h.Pop()
	}
	t3 := time.Now()
	fmt.Printf("heap sort:%d", t3.Sub(t2))
}
