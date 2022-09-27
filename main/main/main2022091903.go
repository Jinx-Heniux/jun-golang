package main

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
