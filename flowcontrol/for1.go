package flowcontrol

import "fmt"

func For1Example1() {
	s := "abc"

	for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
		fmt.Printf("s[%d]=%c(%d)\n", i, s[i], s[i])
	}
	fmt.Println()

	n := len(s) - 1
	for n >= 0 { // 替代 while (n > 0) {}
		fmt.Printf("s[%d]=%c(%d)\n", n, s[n], s[n]) // 替代 for (; n > 0;) {}
		n--
	}
	fmt.Println()

	// for true {}
	i := 0
	for { // 替代 while (true) {}
		if i == len(s) {
			break
		}
		fmt.Printf("s[%d]=%c(%d)\n", i, s[i], s[i]) // 替代 for (;;) {}
		i++
	}
}

/*
s[0]=a(97)
s[1]=b(98)
s[2]=c(99)

s[2]=c(99)
s[1]=b(98)
s[0]=a(97)

s[0]=a(97)
s[1]=b(98)
s[2]=c(99)
*/

func lengthFor1Example2(s string) int {
	fmt.Println("call length.")
	return len(s)
}

func For1Example2() {
	s := "abcd"

	for i, n := 0, lengthFor1Example2(s); i < n; i++ { // 避免多次调用 length 函数。
		// println(i, s[i])
		fmt.Println(i, s[i])
		fmt.Printf("i = %d; value = %c (%d) (%#v) \n", i, s[i], s[i], s[i])
		fmt.Println(i, s[i])
		// println(i, s[i])
	}
}

/*
call length.
0 97
i = 0; value = a (97) (0x61)
0 97
1 98
i = 1; value = b (98) (0x62)
1 98
2 99
i = 2; value = c (99) (0x63)
2 99
3 100
i = 3; value = d (100) (0x64)
3 100
*/

func For1Example3() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	fmt.Printf("numbers: %v\n\n", numbers)

	fmt.Println("a =", a)
	/* for 循环 */
	for a := 0; a < 10; a++ { // 这里的a是在for循环内的局部变量
		fmt.Printf("a 的值为: %d\n", a)
	}

	fmt.Println()

	fmt.Println("a =", a)
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}
	fmt.Println("a =", a)
	fmt.Println()

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

/*
numbers: [1 2 3 5 0 0]

a = 0
a 的值为: 0
a 的值为: 1
a 的值为: 2
a 的值为: 3
a 的值为: 4
a 的值为: 5
a 的值为: 6
a 的值为: 7
a 的值为: 8
a 的值为: 9

a = 0
a 的值为: 1
a 的值为: 2
a 的值为: 3
a 的值为: 4
a 的值为: 5
a 的值为: 6
a 的值为: 7
a 的值为: 8
a 的值为: 9
a 的值为: 10
a 的值为: 11
a 的值为: 12
a 的值为: 13
a 的值为: 14
a 的值为: 15
a = 15

第 0 位 x 的值 = 1
第 1 位 x 的值 = 2
第 2 位 x 的值 = 3
第 3 位 x 的值 = 5
第 4 位 x 的值 = 0
第 5 位 x 的值 = 0
*/

// 使用循环嵌套来输出 2 到 100 间的素数
func For1PrimeNumber1() {
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

/*
2  是素数
3  是素数
5  是素数
7  是素数
11  是素数
13  是素数
17  是素数
19  是素数
23  是素数
29  是素数
31  是素数
37  是素数
41  是素数
43  是素数
47  是素数
53  是素数
59  是素数
61  是素数
67  是素数
71  是素数
73  是素数
79  是素数
83  是素数
89  是素数
97  是素数
*/
