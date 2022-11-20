package flowcontrol

import "fmt"

func Range1Example1() {
	s := "abc"

	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		fmt.Println(s[i])
	}
	fmt.Println()

	// 忽略 index。
	for _, c := range s {
		fmt.Println(c)
	}
	fmt.Println()

	for i, c := range s {
		fmt.Printf("i: %d | %T || c: %d | %T\n", i, i, c, c)
	}
	fmt.Println()

	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2}
	// 返回 (key, value)。
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
97
98
99

97
98
99

i: 0 | int || c: 97 | int32
i: 1 | int || c: 98 | int32
i: 2 | int || c: 99 | int32

b 2
a 1
*/

func Range1Example2() {
	a := [3]int{0, 1, 2}
	fmt.Printf("a: (%p) (%T) %v\n", &a, a, a)

	for i, v := range a { // index、value 都是从复制品中取出。
		fmt.Printf("a in for: (%p) (%T) %v\n", &a, a, a)

		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			// fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
			fmt.Printf("a in if: (%p) (%T) %v\n", &a, a, a)
		}

		fmt.Printf("v in for: (%p) (%T) %v\n", &v, v, v)
		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。

	}

	// fmt.Println(a) // 输出 [100, 101, 102]。
	fmt.Printf("a: (%p) (%T) %v\n", &a, a, a)
}

/*
a: (0xc00001a2d0) ([3]int) [0 1 2]
a in for: (0xc00001a2d0) ([3]int) [0 1 2]
a in if: (0xc00001a2d0) ([3]int) [0 999 999]
v in for: (0xc00001c110) (int) 0
a in for: (0xc00001a2d0) ([3]int) [100 999 999]
v in for: (0xc00001c110) (int) 1
a in for: (0xc00001a2d0) ([3]int) [100 101 999]
v in for: (0xc00001c110) (int) 2
a: (0xc00001a2d0) ([3]int) [100 101 102]
*/

func Range1Example3() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("s: (%p) (%T) %v (%p) %v %v \n", &s, s, s, s, len(s), cap(s))

	for i, v := range s { // 复制 struct slice { pointer, len, cap }。

		if i == 0 {
			s = s[:3] // 对 slice 的修改，不会影响 range。
			fmt.Printf("s: (%p) (%T) %v (%p) %v %v \n", &s, s, s, s, len(s), cap(s))
			s[2] = 100 // 对底层数据的修改。
		}

		fmt.Printf("i: %d || v: (%p) (%T) %v\n", i, &v, v, v)
	}
}

/*
s: (0xc0000ac018) ([]int) [1 2 3 4 5] (0xc0000b2030) 5 5
s: (0xc0000ac018) ([]int) [1 2 3] (0xc0000b2030) 3 5
i: 0 || v: (0xc0000be040) (int) 1
i: 1 || v: (0xc0000be040) (int) 2
i: 2 || v: (0xc0000be040) (int) 100
i: 3 || v: (0xc0000be040) (int) 4
i: 4 || v: (0xc0000be040) (int) 5
*/
