package main

//////////////////////////////////////////////////
// 异常处理 · Go语言中文文档
// https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BC%82%E5%B8%B8%E5%A4%84%E7%90%86.html
// learn_exception
//////////////////////////////////////////////////

/*
import (
	"github.com/Jinx-Heniux/jun-golang-hello-world/learn_exception"
)

func main() {
	learn_exception.Test()
}
*/

/*
panic error!
*/

//////////////////////////////////////////////////

/*
import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}
*/

/*
send on closed channel
*/

//////////////////////////////////////////////////

/*
import "fmt"

func test() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func main() {
	test()
}
*/

/*
defer panic
*/

//////////////////////////////////////////////////
