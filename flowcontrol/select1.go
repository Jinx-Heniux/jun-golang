package flowcontrol

import (
	"fmt"
	"time"
)

var c1 = make(chan int, 5)
var c2, c3 chan int
var i1, i2 int

func Select1Case1() {

	go func() {
		for i := 1; i < 10; i++ {
			// time.Sleep(time.Millisecond * 2)
			c1 <- i
			fmt.Println(c1)
		}
	}()
	go testSelect1Case1()
}

func testSelect1Case1() {
	time.Sleep(time.Second * 1)
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

var resChan = make(chan int)

func Select1Case2() {

	// do request
	go testSelect1Case2()
	resChan <- 1

	// time.Sleep(time.Second * 1)

}

func testSelect1Case2() {
	select {
	case data := <-resChan:
		doDataSelect1Case2(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}

func doDataSelect1Case2(data int) {
	//...
	fmt.Printf("in doData...%d\n", data)
}
