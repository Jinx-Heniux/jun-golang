package struct2_

import "fmt"

//    go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

//人
type Person struct {
	name string
	sex  string
	age  int
}

type Student1 struct {
	Person
	id   int
	addr string
}

type Student2 struct {
	Person
	id   int
	addr string
	//同名字段
	name string
}

type Student3 struct {
	*Person
	id   int
	addr string
}

func AnonymousFieldExample1() {
	// 初始化
	s1 := Student1{Person{"5lmh", "man", 20}, 1, "bj"}
	fmt.Println(s1)

	s2 := Student1{Person: Person{"5lmh", "man", 20}}
	fmt.Println(s2)

	s3 := Student1{Person: Person{name: "5lmh"}}
	fmt.Println(s3)
}

/*
{{5lmh man 20} 1 bj}
{{5lmh man 20} 0 }
{{5lmh  0} 0 }
*/

func AnonymousFieldExample2() {
	var s Student2 = Student2{id: 1, addr: "abc"}
	// 给自己字段赋值了
	s.name = "5lmh"
	fmt.Println(s)

	// 若给父类同名字段赋值，如下
	s.Person.name = "枯藤"
	fmt.Println(s)
}

/*
{{  0} 1 abc 5lmh}
{{枯藤  0} 1 abc 5lmh}
*/

func AnonymousFieldExample3() {
	s1 := Student3{&Person{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)
}

/*
{0xc0000a0150 1 bj}
5lmh
5lmh
*/
