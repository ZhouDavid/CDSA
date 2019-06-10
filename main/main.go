package main

import (
	"GoSTL/list"
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

	// test list
	l1 := list.New()
	for i := 0; i < 10; i++ {
		l1.PushBack(rand.Float64())
	}
	l1.Insert(3, 123123.)
	// sort.Sort(l1)
	for p := l1.Front(); p != l1.End(); p = p.Next() {
		fmt.Printf("%v ", p.Value)
	}
	fmt.Println()
	sort.Sort(l1)
	for p := l1.Front(); p != l1.End(); p = p.Next() {
		fmt.Printf("%v ", p.Value)
	}
}
