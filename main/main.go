package main

import (
	"GoSTL/vector"
	"fmt"
)

func main() {
	arr := vector.NewVector(10, 0)
	fmt.Printf("before insert:%v\n", arr)
	if err := arr.Insert(0, 1); err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("after insert:%v\n", arr)
	arr.Remove(0)
	fmt.Printf("after remove:%v\n", arr)

}
