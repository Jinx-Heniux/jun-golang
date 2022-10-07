package struct_

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func StructExample1() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for i, stu := range stus {
		// fmt.Printf("stu.name: %v | stu: %v | stus[%d]: %v\n", stu.name, stu, i, stus[i])
		// fmt.Printf("addr stu: %p | addr of stus[%d]: %p\n", &stu, i, &stus[i])
		fmt.Printf("stus: %v | addr: %p\n", stus, &stus)
		fmt.Printf("\tstu: %v | addr: %p\n", stu, &stu)
		fmt.Printf("\tstus[%d]: %v | addr: %p\n", i, stus[i], &stus[i])
		fmt.Println()
		// m[stu.name] = &stu
		// 注意：stu的地址是不变的。 应该使用&stus[i]
		m[stu.name] = &stus[i]
	}
	fmt.Println()
	for k, v := range m {
		fmt.Printf("%#v => %#v (%T) (%p) \n", k, v, v, v)
	}
}

/*
stus: [{pprof.cn 18} {测试 23} {博客 28}] | addr: 0xc00000c030
        stu: {pprof.cn 18} | addr: 0xc00000c048
        stus[0]: {pprof.cn 18} | addr: 0xc000064050

stus: [{pprof.cn 18} {测试 23} {博客 28}] | addr: 0xc00000c030
        stu: {测试 23} | addr: 0xc00000c048
        stus[1]: {测试 23} | addr: 0xc000064068

stus: [{pprof.cn 18} {测试 23} {博客 28}] | addr: 0xc00000c030
        stu: {博客 28} | addr: 0xc00000c048
        stus[2]: {博客 28} | addr: 0xc000064080


"测试" => &main.student{name:"博客", age:28} (*main.student) (0xc00000c048)
"博客" => &main.student{name:"博客", age:28} (*main.student) (0xc00000c048)
"pprof.cn" => &main.student{name:"博客", age:28} (*main.student) (0xc00000c048)
*/
