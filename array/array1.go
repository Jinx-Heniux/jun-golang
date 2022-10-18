package array

import (
	"fmt"
	"math/rand"
	"time"
)

// 全局
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "Hello", 4: "Tom"}
var arr [5]int32 // [0 0 0 0 0]

func ArrayExample1() {

	// 局部
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3}
	c := [5]int{0: 100, 4: 100}
	d := [...]struct {
		name string
		age  uint8
	}{
		// {name: "user1", age: 10},
		// {name: "user2", age: 20},
		{"user1", 10},
		{"user2", 20},
	}

	fmt.Println(arr0, arr1, arr2, str, arr)
	fmt.Println(a, b, c, d)

}

/*
[1 2 3 0 0] [1 2 3 4 5] [1 2 3 4 5 6] [   Hello Tom] [0 0 0 0 0]
[1 2 0] [1 2 3] [100 0 0 0 100] [{user1 10} {user2 20}]
*/

var arr20 [5][3]int
var arr21 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func ArrayExample2() {

	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}} // 第 2 纬度不能用 "..."。
	// fmt.Println(arr20, arr21)
	// fmt.Println(a, b)

	fmt.Println(arr20)
	fmt.Println(arr21)
	fmt.Println(a)
	fmt.Println(b)

}

/*
[[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]] [[1 2 3] [7 8 9]]
[[1 2 3] [4 5 6]] [[1 2] [2 3] [3 4] [4 5]]
*/

/*
[[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]]
[[1 2 3] [7 8 9]]
[[1 2 3] [4 5 6]]
[[1 2] [2 3] [3 4] [4 5]]
*/

func test(x [2]int) {
	fmt.Printf("x -> %#v || ", x)
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
	fmt.Printf("x -> %#v || ", x)
	fmt.Printf("x: %p\n", &x)
}

func ArrayExample3() {

	a := [2]int{}
	fmt.Printf("a -> %#v || ", a)
	fmt.Printf("a: %#p\n", &a)
	test(a)
	fmt.Printf("a -> %#v || ", a)
	fmt.Printf("a: %p\n", &a)
	// fmt.Println(a)

	fmt.Println(len(a), cap(a)) // 内置函数 len 和 cap 都返回数组长度 (元素数量)。

}

/*
a -> [2]int{0, 0} || a: c000128000
x -> [2]int{0, 0} || x: 0xc000128040
x -> [2]int{0, 1000} || x: 0xc000128040
a -> [2]int{0, 0} || a: 0xc000128000
2 2
*/

func TraverseArray1() {

	var arr [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

	for k1, v1 := range arr {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d	", k1, k2, v2)
		}
		fmt.Println()
	}

}

/*
(0,0)=1 (0,1)=2 (0,2)=3
(1,0)=7 (1,1)=8 (1,2)=9
*/

func printArr(arr *[5]int) {
	arr[0] = 10
	for k, v := range arr {
		fmt.Printf("%d -> %d\n", k, v)
	}
}

func ModifyArray1() {
	// arr1 := [5]int{}
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)

	arr2 := [...]int{1, 2, 3, 4, 5}
	printArr(&arr2)
	fmt.Println(arr2)
}

/*
0 -> 10
1 -> 0
2 -> 0
3 -> 0
4 -> 0
[10 0 0 0 0]
0 -> 10
1 -> 2
2 -> 3
3 -> 4
4 -> 5
[10 2 3 4 5]
*/

func sumArr(arr [10]int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func SumArray1() {
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	// rand.Seed(time.Now().Unix())
	t := time.Now()
	fmt.Println(t)
	i := t.Unix()
	fmt.Println(i)
	rand.Seed(i)

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
		fmt.Println(b[i])
	}
	sum := sumArr(b)
	fmt.Printf("sum=%d\n", sum)
}

// 求元素和，是给定的值
func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func ArrayExample4() {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}
