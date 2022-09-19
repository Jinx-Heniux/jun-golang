package main

//////////////////////////////////////////////////
// 异常处理 · Go语言中文文档
// https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BC%82%E5%B8%B8%E5%A4%84%E7%90%86.html
// learn_exception
//////////////////////////////////////////////////

/*
import (
	"fmt"
	"time"

	"github.com/Jinx-Heniux/jun-golang-hello-world/learn_exception"
)

func main() {


	fmt.Println("##############################")
	learn_exception.Test1()
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	learn_exception.Test2()
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	learn_exception.Test3()
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	learn_exception.Test4()
	time.Sleep(time.Second * 2)

}
*/

/*
import "fmt"

func test() {
	// defer func() {
	// 	fmt.Println(recover()) //有效
	// }()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()
	defer func() {
		fmt.Println(recover()) //有效
	}()

	panic("test panic")
}

func main() {
	test()
}
*/
