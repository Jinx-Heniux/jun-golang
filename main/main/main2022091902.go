package main

/*
import (
	"fmt"

	"github.com/Jinx-Heniux/jun-golang-hello-world/learn_exception"
)

func main() {

	// fmt.Println("##############################")
	// learn_exception.Test4()

}
*/

/*
import "fmt"

// compared to learn_exception.Test4()

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
