package main

//////////////////////////////////////////////////
// 异常处理 · Go语言中文文档
// https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BC%82%E5%B8%B8%E5%A4%84%E7%90%86.html
// learn_exception
//////////////////////////////////////////////////

/*
import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := div(10, 0); err {
	case nil:
		println("in case nil: ")
		println(z)
	case ErrDivByZero:
		println("in case ErrDivByZero: ")
		println(z)
		panic(err)
	}
}
*/

/*
in case ErrDivByZero:
0
division by zero
*/
