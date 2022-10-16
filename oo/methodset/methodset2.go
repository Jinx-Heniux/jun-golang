package methodset

import (
	"fmt"
)

type T2 struct {
	int
}

func (t T2) testT() {
	fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
	fmt.Printf("in testT: %v | %p\n", t, &t)
}

func (t *T2) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
	fmt.Printf("in testT: %v | %p | %p\n", t, &t, t)
}

func MethodSetExample2() {
	t1 := T2{1}
	fmt.Printf("t1: %v | %p\n", t1, &t1)
	t1.testT()
	t1.testP()
	fmt.Println()

	t2 := &t1
	fmt.Printf("t2: %v | %p | %p\n", t2, &t2, t2)
	t2.testT()
	t2.testP()
}

/*
t1: {1} | 0xc00001c0f8
类型 *T 方法集包含全部 receiver T 方法。
in testT: {1} | 0xc00001c108
类型 *T 方法集包含全部 receiver *T 方法。
in testT: &{1} | 0xc00000e030 | 0xc00001c0f8

t2: &{1} | 0xc00000e038 | 0xc00001c0f8
类型 *T 方法集包含全部 receiver T 方法。
in testT: {1} | 0xc00001c130
类型 *T 方法集包含全部 receiver *T 方法。
in testT: &{1} | 0xc00000e040 | 0xc00001c0f8
*/
