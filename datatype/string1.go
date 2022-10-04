package datatype

import (
	"fmt"
)

// 遍历字符串
func TraversalString() {

	s := "pprof.cn博客"
	fmt.Printf("s -> %v(%s)(%T)(%d)\n", s, s, s, len(s))
	// 字符串由字符组成，一个中文字等于3个字符
	fmt.Println()

	for i := 0; i < len(s); i++ { //byte
		// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
		fmt.Printf("%v(%c)(%T) ", s[i], s[i], s[i])
	}
	fmt.Println()
	fmt.Println()

	rune_nums := 0
	//for _, r := range s { //rune
	for i, r := range s {
		// rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
		fmt.Printf("%d -> %v(%c)(%T) \n", i, r, r, r)
		rune_nums += 1
	}
	fmt.Println()
	fmt.Printf("numbers of runes: %d\n", rune_nums)
}

/*
s -> pprof.cn博客(pprof.cn博客)(string)(14)

112(p)(uint8) 112(p)(uint8) 114(r)(uint8) 111(o)(uint8) 102(f)(uint8) 46(.)(uint8) 99(c)(uint8) 110(n)(uint8) 229(å)(uint8) 141()(uint8) 154()(uint8) 229(å)(uint8) 174(®)(uint8) 162(¢)(uint8)

0 -> 112(p)(int32)
1 -> 112(p)(int32)
2 -> 114(r)(int32)
3 -> 111(o)(int32)
4 -> 102(f)(int32)
5 -> 46(.)(int32)
6 -> 99(c)(int32)
7 -> 110(n)(int32)
8 -> 21338(博)(int32)
11 -> 23458(客)(int32)

numbers of runes: 10
*/

// 修改字符串
// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
// 无论哪种转换，都会重新分配内存，并复制字节数组。
func ChangeString() {

	s1 := "hello"

	// 强制类型转换
	byteS1 := []byte(s1)
	fmt.Printf("byteS1: %v (%c) (%#v) (%T) (%d)\n", byteS1, byteS1, byteS1, byteS1, len(byteS1))
	fmt.Println()

	byteS1[0] = 'H'
	fmt.Println(string(byteS1))
	fmt.Println()

	s2 := "博客"

	runeS2 := []rune(s2)
	fmt.Printf("runeS2: %v (%c) (%#v) (%T) (%d)\n", runeS2, runeS2, runeS2, runeS2, len(runeS2))
	fmt.Println()

	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

/*
byteS1: [104 101 108 108 111] ([h e l l o]) ([]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}) ([]uint8) (5)

Hello
runeS2: [21338 23458] ([博 客]) ([]int32{21338, 23458}) ([]int32) (2)

狗客
*/
