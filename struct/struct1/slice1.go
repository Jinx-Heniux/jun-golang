package struct1

import "fmt"

type student1 struct {
	id   int
	name string
	age  int
}

func demo(ce []student1) {
	//切片是引用传递，是可以改变值的
	fmt.Printf("demo: %p | %p | %v\n", &ce, ce, ce)
	ce[1].age = 999
	ce = append(ce, student1{3, "xiaowang", 56})
	fmt.Printf("demo: %p | %p | %v\n", &ce, ce, ce)
	// return ce

	// 注意：这里实际上还是传递的[]student1的副本，
	// 只不过副本中包含了指向底层数组的引用，所以可以修改，
	// 但当添加新元素时，出现问题。
	// 在demo中使用append后，底层数组改变。
}
func SliceExample1() {
	/*
		var ce []student1 //定义一个切片类型的结构体
		ce = []student1{
			student1{1, "xiaoming", 22},
			student1{2, "xiaozhang", 33},
		}
	*/

	var ce = []student1{
		{1, "xiaoming", 22},
		{2, "xiaozhang", 33},
	}
	// fmt.Println(ce)
	fmt.Printf("%p | %p | %v\n", &ce, ce, ce)
	demo(ce)
	// fmt.Println(ce)
	fmt.Printf("%p | %p | %v\n", &ce, ce, ce)
}

/*
0xc0000ac018 | 0xc0000ba040 | [{1 xiaoming 22} {2 xiaozhang 33}]
demo: 0xc0000ac060 | 0xc0000ba040 | [{1 xiaoming 22} {2 xiaozhang 33}]
demo: 0xc0000ac060 | 0xc0000c8000 | [{1 xiaoming 22} {2 xiaozhang 999} {3 xiaowang 56}]
0xc0000ac018 | 0xc0000ba040 | [{1 xiaoming 22} {2 xiaozhang 999}]
*/
