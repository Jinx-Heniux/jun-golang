package slice2

import (
	"fmt"
	"unsafe"
)

func Slice2Array1() {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

// arrayA : 0xc000138000 , [100 200]
// arrayB : 0xc000138010 , [100 200]
// func Array : 0xc000138050 , [100 200]
// 三个内存地址都不同，这也就验证了 Go 中数组赋值和函数传参都是值复制的。

func Slice2Array2() {
	arrayA := [2]int{100, 200}
	fmt.Printf("(%T)(%d)(%d)(%p)\n", arrayA, len(arrayA), cap(arrayA), &arrayA)
	fmt.Printf("(%T)(%d)(%d)(%p)\n", arrayA, len(arrayA), cap(arrayA), &arrayA[0])

	fmt.Println()
	testArrayAPoint(&arrayA) // 1.传数组指针
	fmt.Println()

	arrayB := arrayA[:]
	testArrayBPoint(&arrayB) // 2.传切片
	fmt.Println()

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)
	fmt.Printf("arrayB[0] : %p , %v\n", &arrayB[0], arrayB[0])
}

func testArrayAPoint(x *[2]int) {
	fmt.Printf("func ArrayA : %p , %v\n", x, *x)
	(*x)[1] += 100
	fmt.Printf("func ArrayA : %p , %v\n", x, *x)
}

func testArrayBPoint(x *[]int) {
	fmt.Printf("func ArrayB : %p , %v\n", x, *x)
	(*x)[1] += 100
	fmt.Printf("func ArrayB : %p , %v\n", x, *x)
}

/*
([2]int)(2)(2)(0xc00001c110)
([2]int)(2)(2)(0xc00001c110)

func ArrayA : 0xc00001c110 , [100 200]
func ArrayA : 0xc00001c110 , [100 300]

func ArrayB : 0xc00000c030 , [100 300]
func ArrayB : 0xc00000c030 , [100 400]

arrayA : 0xc00001c110 , [100 400]
arrayB : 0xc00000c030 , [100 400]
arrayB[0] : 0xc00001c110 , 100
*/

func Slice2Pointer1() {
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0]) // 从 slice 中得到一块内存地址
	fmt.Printf("&s:%p | &arr:%p | %p\n", &s, ptr, s)
	// &s:0xc00000c030 | &arr:0xc00007e000 | 0xc00007e000
}

func Slice2Pointer2() {
	var ptr unsafe.Pointer
	fmt.Printf("ptr: %T | %v | %p | %p\n", ptr, ptr, &ptr, ptr)
	fmt.Println()

	var s1 = struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{ptr, 2, 2}
	fmt.Printf("s1: %T | %v | %p | %p\n", s1, s1, &s1, &s1.addr)
	fmt.Println()

	s := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Printf("s: %T | %v | %p | %p", s, s, &s, s)

}
