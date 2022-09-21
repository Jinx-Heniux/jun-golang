package learn_exception

import "fmt"

func Test1() {
	defer func() {
		println("Test1(): defer 0")
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
		println("Test1(): defer 0")
	}()

	panic("panic error in Test1()!")
}

func Test2() {
	defer func() {
		println("Test2(): defer 0")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		println("Test2(): defer 0")
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

func Test3() {
	defer func() {
		println("Test3(): defer 0")
		fmt.Println(recover())
		println("Test3(): defer 0")
	}()

	defer func() {
		println("Test3(): defer 1")
		panic("Test3(): defer panic")
	}()

	panic("Test3(): test panic")
}

func Test4() {
	defer func() {
		println("Test4(): defer 0")
		fmt.Println(recover()) //有效
		println("Test4(): defer 0")
	}()
	defer fmt.Println(recover())
	defer func() {
		println("Test4(): defer 1")
		fmt.Println(recover()) //有效
		println("Test4(): defer 1")
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		println("Test4(): defer 2")
		func() {
			println("Test4(): defer inner")
			recover() //无效！
			println("Test4(): defer 2")
		}()
		println("Test4(): defer 2")
	}()
	defer func() {
		println("Test4(): defer 3")
		fmt.Println(recover()) //有效
		println("Test4(): defer 3")
	}()

	panic("Test4(): test panic")
}

func except() {
	fmt.Println(recover())
}

func Test5() {
	defer except()
	panic("test panic")
}

func Test6(x, y int) {
	var z int
	fmt.Println(z)
	func() {
		defer func() {
			if recover() != nil {
				fmt.Println(z)
				z = 100
				fmt.Println(z)
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
