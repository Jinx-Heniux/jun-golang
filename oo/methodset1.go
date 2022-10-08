package oo

import (
	"fmt"
)

type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
	fmt.Printf("in testT: %v | %p\n", t, &t)
}

func MethodSetExample1() {
	t1 := T{1}
	fmt.Printf("t1: %v | %p\n", t1, &t1)
	t1.test()
	fmt.Println()

	t2 := &t1
	fmt.Printf("t2: %v | %p | %p\n", t2, &t2, t2)
	t2.test()
}

/*
t1: {1} | 0xc0000ba000
类型 T 方法集包含全部 receiver T 方法。
in testT: {1} | 0xc0000ba020

t2: &{1} | 0xc0000b4020 | 0xc0000ba000
类型 T 方法集包含全部 receiver T 方法。
in testT: {1} | 0xc0000ba030
*/
