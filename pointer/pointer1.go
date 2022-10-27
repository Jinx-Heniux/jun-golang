package pointer

import (
	"fmt"
)

func Pointer1Example1() {
	a := 10
	b := &a
	fmt.Printf("a: %d, &a: %p\n", a, &a)
	fmt.Printf("b: %#v | %v | %p, &b: %p, *b: %d", b, b, b, &b, *b)
}

/*
a: 10, &a: 0xc0000ba000
b: (*int)(0xc0000ba000) | 0xc0000ba000, &b: 0xc0000b4018, *b: 10
*/

func Pointer1Example2() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

/*
type of b:*int
type of c:int
value of c:10
*/

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func Pointer1Example3() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}

// ///////////////////////////////

func Pointer1Nil1() {

	var p *string
	fmt.Printf("address of p: %v\n", &p)
	fmt.Printf("type of p: %T\n", p)
	fmt.Printf("value of p: %v\n", p)
	if p != nil {
		fmt.Println("not null")
	} else {
		fmt.Println("null")
	}
	//fmt.Printf("value of p: %s\n", *p)
	// panic: runtime error: invalid memory address or nil pointer dereference
	d := "Hello"
	p = &d
	fmt.Println("++++++")
	fmt.Printf("address of p: %v\n", &p)
	fmt.Printf("type of p: %T\n", p)
	fmt.Printf("value of p: %v\n", p)
	fmt.Printf("value of *p: %s\n", *p)

	if p != nil {
		fmt.Println("not null")
	} else {
		fmt.Println("null")
	}

}

/*
address of p: 0xc0000bc018
type of p: *string
value of p: <nil>
null
++++++
address of p: 0xc0000bc018
type of p: *string
value of p: 0xc000096230
value of *p: Hello
not null
*/

func Pointer1Example4() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}

/*
执行上面的代码会引发panic，为什么呢？
在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，
还要为它分配内存空间，否则我们的值就没办法存储。
*/

func Pointer1New1() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a) // 10
}

func Pointer1New2() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

/*
new函数不太常用，使用new函数得到的是一个类型的指针，
并且该指针对应的值为该类型的零值。
*/

func Pointer1Make1() {
	var b map[string]int         // 只是声明变量b是一个map类型的变量
	b = make(map[string]int, 10) // 使用make函数进行初始化操作之后，才能对其进行键值对赋值
	b["测试"] = 100
	fmt.Println(b) // map[测试:100]
}

func Pointer1Example5() {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}

/*
0xc0000b8000
20
*/
