package main

import (
	"fmt"
	"time"

	"github.com/Jinx-Heniux/jun-golang/exception"
)

func main() {

	fmt.Println("##############################")
	exception.Test1()
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	exception.Test2()
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	exception.Test3()
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	exception.Test5()
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	exception.Test6(2, 1)
	time.Sleep(time.Second * 2)

	fmt.Println("##############################")
	exception.Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})

}
