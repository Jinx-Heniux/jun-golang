package slice2

import (
	"fmt"
)

var slice = make([]int, 0)
var slice0 []int = make([]int, 10)
var slice1 = make([]int, 10)
var slice2 = make([]int, 10, 10)

func Slice2Make1() {
	fmt.Printf("(%T)(%d)(%d) ", slice, len(slice), cap(slice))
	fmt.Printf("make全局slice : %v\n", slice)
	fmt.Printf("(%T)(%d)(%d) ", slice0, len(slice0), cap(slice0))
	fmt.Printf("make全局slice0 : %v\n", slice0)
	fmt.Printf("(%T)(%d)(%d) ", slice1, len(slice1), cap(slice1))
	fmt.Printf("make全局slice1 : %v\n", slice1)
	fmt.Printf("(%T)(%d)(%d) ", slice2, len(slice2), cap(slice2))
	fmt.Printf("make全局slice2 : %v\n", slice2)
	fmt.Println("--------------------------------------")
	slice3 := make([]int, 10)
	slice4 := make([]int, 10)
	slice5 := make([]int, 10, 10)
	slice6 := make([]int, 0, 10)
	slice7 := make([]int, 0)
	fmt.Printf("(%T)(%d)(%d) ", slice3, len(slice3), cap(slice3))
	fmt.Printf("make局部slice3 : %v\n", slice3)
	fmt.Printf("make局部slice4 : %v\n", slice4)
	fmt.Printf("make局部slice5 : %v\n", slice5)
	fmt.Printf("(%T)(%d)(%d) ", slice6, len(slice6), cap(slice6))
	fmt.Printf("make局部slice6 : %v\n", slice6)
	fmt.Printf("(%T)(%d)(%d) ", slice7, len(slice7), cap(slice7))
	fmt.Printf("make局部slice7 : %v\n", slice7)
}

/*
([]int)(0)(0) make全局slice : []
([]int)(10)(10) make全局slice0 : [0 0 0 0 0 0 0 0 0 0]
([]int)(10)(10) make全局slice1 : [0 0 0 0 0 0 0 0 0 0]
([]int)(10)(10) make全局slice2 : [0 0 0 0 0 0 0 0 0 0]
--------------------------------------
([]int)(10)(10) make局部slice3 : [0 0 0 0 0 0 0 0 0 0]
make局部slice4 : [0 0 0 0 0 0 0 0 0 0]
make局部slice5 : [0 0 0 0 0 0 0 0 0 0]
([]int)(0)(10) make局部slice6 : []
([]int)(0)(0) make局部slice7 : []
*/
