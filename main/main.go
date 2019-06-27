package main

import (
	"GoSTL/util"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Array []float64

func (arr *Array) Push(x interface{}) {
	*arr = append(*arr, x.(float64))
}
func (arr *Array) Pop() interface{} {
	n := arr.Len()
	v := (*arr)[n-1]
	*arr = (*arr)[:n-1]
	return v
}
func (arr *Array) Swap(i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
func (arr *Array) Less(i, j int) bool {
	return (*arr)[i] < (*arr)[j]
}
func (arr *Array) Len() int {
	return len((*arr))
}
func (arr *Array) Get(i int) interface{} {
	return (*arr)[i]
}
func (arr *Array) Set(i int, value interface{}) {
	(*arr)[i] = value.(float64)
}

func main() {
	// rand.Seed(int64(time.Now().UnixNano()))
	arrLen := 10000000 * 4
	t0 := time.Now()
	arr1 := make(Array, arrLen)
	arr2 := make(Array, arrLen)
	for i := 0; i < arrLen; i++ {
		arr1[i] = rand.Float64()
		arr2[i] = arr1[i]
	}

	t1 := time.Now()
	fmt.Printf("%v\n", t1.Sub(t0))
	t2 := time.Now()
	util.MultiThreadQuickSort(&arr2)
	t3 := time.Now()
	fmt.Printf("mutithread quickSort with insertion:%v\n\n", t3.Sub(t2))
	t3 = time.Now()
	util.MultiThreadQuickSort(&arr1)
	t4 := time.Now()
	fmt.Printf("multithread quickSort arr:%v\n\n", t4.Sub(t3))
	fmt.Printf("%v\n\n", sort.IsSorted(&arr1))
	fmt.Printf("%v\n\n", sort.IsSorted(&arr2))
}
