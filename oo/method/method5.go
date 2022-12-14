package method

import "fmt"

//定义一个结构体
type T5 struct {
	name string
}

func (t T5) method1() {
	t.name = "new name1"
}

func (t *T5) method2() {
	t.name = "new name2"
}

func MethodExample5() {

	t := T5{"old name"}

	fmt.Println("method1 调用前 ", t.name)
	t.method1()
	fmt.Println("method1 调用后 ", t.name)

	fmt.Println("method2 调用前 ", t.name)
	t.method2()
	fmt.Println("method2 调用后 ", t.name)
}

/*
method1 调用前  old name
method1 调用后  old name
method2 调用前  old name
method2 调用后  new name2
*/
