package oo

import (
	"fmt"
)

type S3 struct {
	T3
}

type T3 struct {
	int
}

func (t T3) testT() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
	fmt.Printf("in testT: %v | %p\n", t, &t)
}

func (t *T3) testP() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 *T 方法")
	fmt.Printf("in testT: %v | %p | %p\n", t, &t, t)
}

func MethodSetExample3() {
	s1 := S3{T3{1}}
	fmt.Printf("s1: %v | %p\n", s1, &s1)
	s1.testT()
	s1.testP()
	fmt.Println()

	s2 := &s1
	fmt.Printf("s2: %v | %p | %p\n", s2, &s2, s2)
	s2.testT()
	s2.testP()
}

/*
s1: {{1}} | 0xc0000ba000
如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
in testT: {1} | 0xc0000ba020
s2: &{{1}} | 0xc0000b4020 | 0xc0000ba000
如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
in testT: {1} | 0xc0000ba038
*/
