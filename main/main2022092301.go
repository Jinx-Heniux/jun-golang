package main

//////////////////////////////////////////////////
/*
3、Select · 语雀

https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BC%82%E5%B8%B8%E5%A4%84%E7%90%86.html

learn_select/select.go
*/
//////////////////////////////////////////////////

import (
	"fmt"

	"github.com/Jinx-Heniux/jun-golang/learn_select"
)

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	learn_select.Fibonacci(c, quit)
}
