package struct1

import "fmt"

// 修改demo的参数为指针类型
func demo2(ce *[]student1) {
	fmt.Printf("demo 1: %v, %d, %d, %p, %p\n", ce, len(*ce), len(*ce), ce, &(*ce)[0])
	//切片是引用传递，是可以改变值的
	(*ce)[1].age = 999
	*ce = append(*ce, student1{3, "xiaowang", 56})
	fmt.Printf("demo 2: %v, %d, %d, %p, %p\n", ce, len(*ce), len(*ce), ce, &(*ce)[0])
}
func SliceExample2() {
	// var ce []student //定义一个切片类型的结构体
	ce := []student1{
		{1, "xiaoming", 22},
		{2, "xiaozhang", 33},
	}
	// fmt.Println(ce)
	fmt.Printf("main 1: %v, %d, %d, %p, %p\n", ce, len(ce), len(ce), &ce, &ce[0])
	demo2(&ce)
	// fmt.Println(ce)
	fmt.Printf("main 2: %v, %d, %d, %p, %p\n", ce, len(ce), len(ce), &ce, &ce[0])
}

/*
main 1: [{1 xiaoming 22} {2 xiaozhang 33}], 2, 2, 0xc0000ac018, 0xc0000ba040
demo 1: &[{1 xiaoming 22} {2 xiaozhang 33}], 2, 2, 0xc0000ac018, 0xc0000ba040
demo 2: &[{1 xiaoming 22} {2 xiaozhang 999} {3 xiaowang 56}], 3, 3, 0xc0000ac018, 0xc0000c8000
main 2: [{1 xiaoming 22} {2 xiaozhang 999} {3 xiaowang 56}], 3, 3, 0xc0000ac018, 0xc0000c8000
*/
