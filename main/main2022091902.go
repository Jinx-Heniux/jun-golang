package main

/*
import (
	"fmt"

	"github.com/Jinx-Heniux/jun-golang/exception"
)

func main() {

	fmt.Println("##############################")
	exception.Test4()

}
*/

/*
import "fmt"

// compared to exception.Test4()

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
