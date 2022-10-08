package oo

import (
	"fmt"
)

type S4 struct {
	*T4
}

type T4 struct {
	int
}

func (t T4) testT() {
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法")
	fmt.Printf("in testT: %v | %p\n", t, &t)
}
func (t *T4) testP() {
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法")
	fmt.Printf("in testP: %v | %p | %p\n", t, &t, t)
}

func (s S4) testT() {
	fmt.Println("如类型 S 包含匿名字段 *T，则 ......")
	fmt.Printf("in testT: %v | %p\n", s, &s)
}

func MethodSetExample4() {
	s1 := S4{&T4{1}}
	fmt.Printf("s1: %v | %p\n", s1, &s1)
	s1.testT()
	s1.testP()
	s1.T4.testT()
	s1.T4.testP()
	fmt.Println()

	s2 := &s1
	fmt.Printf("s2: %v | %p | %p\n", s2, &s2, s2)
	s2.testT()
	s2.testP()
	s2.T4.testT()
	s2.T4.testP()
}

/*
s1: {0xc00001c0f8} | 0xc00000e028
如类型 S 包含匿名字段 *T，则 ......
in testT: {0xc00001c0f8} | 0xc00000e038
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法
in testP: &{1} | 0xc00000e040 | 0xc00001c0f8
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法
in testT: {1} | 0xc00001c110
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法
in testP: &{1} | 0xc00000e048 | 0xc00001c0f8

s2: &{0xc00001c0f8} | 0xc00000e050 | 0xc00000e028
如类型 S 包含匿名字段 *T，则 ......
in testT: {0xc00001c0f8} | 0xc00000e058
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法
in testP: &{1} | 0xc00000e060 | 0xc00001c0f8
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法
in testT: {1} | 0xc00001c128
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法
in testP: &{1} | 0xc00000e068 | 0xc00001c0f8
*/
