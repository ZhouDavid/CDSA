package main

import (
	"GoSTL/vector"
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// test vector
	arr1 := vector.NewVector(20, 0.)
	for i := 0; i < 5; i++ {
		arr1[i] = rand.Float64()
	}

	fmt.Printf("%v\n", arr1)
	sort.Sort(arr1)
	fmt.Printf("%v", arr1)
	// // test list
	// l1 := linkedlist.New()
	// for i := 0; i < 10; i++ {
	// 	l1.PushBack(i)
	// }
	// // l1.Insert(3, 123123)
	// l1.RemoveRange(1, 10)
	// for p := l1.Front(); p != l1.End(); p = p.Next() {
	// 	fmt.Printf("%d ", p.Value)
	// }
}
