package main

import (
	linkedlist "GoSTL/list"
	"fmt"
)

func main() {
	//// test vector
	// arr1 := vector.NewVector(10, 0)
	// fmt.Printf("before insert:%v\n", arr1)
	// if err := arr1.Insert(0, 1); err != nil {
	// 	fmt.Printf("%v\n", err)
	// }
	// arr2 := vector.NewVector(3, 1)
	// temp := ([]interface{})(*arr2)
	// arr1.InsertRange(1, temp)
	// fmt.Printf("after insert:%v\n", arr1)
	// fmt.Printf("before removal:%v\n", arr1)
	// arr1.RemoveRange(3, 10)
	// fmt.Printf("after removal:%v\n", arr1)
	// fmt.Printf("after remove:%v\n", arr)

	// test list
	l1 := linkedlist.New()
	for i := 0; i < 10; i++ {
		l1.PushBack(i)
	}
	// l1.Insert(0, 123123)
	for p := l1.Front(); p != l1.e; p = p.Next() {
		fmt.Printf("%d ", p.Value)
	}
}
