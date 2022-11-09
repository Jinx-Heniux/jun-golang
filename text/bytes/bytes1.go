package bytes

import (
	"bytes"
	"fmt"
)

func Bytes1Runes1() {
	b := []byte("你好，世界")
	for k, v := range b {
		fmt.Printf("%d: %s\n", k, string(v))
	}
	fmt.Println()

	r := bytes.Runes(b)
	for k, v := range r {
		fmt.Printf("%d: %s\n", k, string(v))
	}
}

/*
0: ä
1: ½
2:
3: å
4: ¥
5: ½
6: ï
7: ¼
8: 
9: ä
10: ¸
11: 
12: ç
13: 
14: 

0: 你
1: 好
2: ，
3: 世
4: 界
*/

func Bytes1Reader1() {
	x := []byte("你好，世界")

	r1 := bytes.NewReader(x)
	d1 := make([]byte, len(x))
	n, _ := r1.Read(d1)
	fmt.Println(n, string(d1))

	r2 := bytes.Reader{}
	r2.Reset(x)
	d2 := make([]byte, len(x))
	n, _ = r2.Read(d2)
	fmt.Println(n, string(d2))
}

/*
15 你好，世界
15 你好，世界
*/

func Bytes1Reader2() {
	x := []byte("你好，世界")
	r1 := bytes.NewReader(x)

	ch, size, _ := r1.ReadRune()
	fmt.Println(size, ch, string(ch))
	_ = r1.UnreadRune() // 相当于，将读取的指针回退一个Rune
	ch, size, _ = r1.ReadRune()
	fmt.Println(size, ch, string(ch))
	_ = r1.UnreadRune()

	by, _ := r1.ReadByte()
	fmt.Println(by, string(by))
	_ = r1.UnreadByte()
	by, _ = r1.ReadByte()
	fmt.Println(by, string(by))
	_ = r1.UnreadByte()

	d1 := make([]byte, 6)
	n, _ := r1.Read(d1)
	fmt.Println(n, d1, string(d1))
	/*
		d2 := make([]byte, 6)
		n, _ = r1.ReadAt(d2, 0)
		fmt.Println(n, d2, string(d2))
	*/
	d2 := make([]byte, 3)
	n, _ = r1.ReadAt(d2, 3)
	fmt.Println(n, d2, string(d2))

	w1 := &bytes.Buffer{}
	_, _ = r1.Seek(0, 0)
	_, _ = r1.WriteTo(w1)
	fmt.Println(w1.String())
}

/*
3 20320 你
3 20320 你
228 ä
228 ä
6 [228 189 160 229 165 189] 你好
3 [229 165 189] 好
你好，世界
*/

func Bytes1Buffer1() {
	a := bytes.NewBufferString("Hello World")
	b := bytes.NewBuffer([]byte("Hello World"))
	c := bytes.Buffer{}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*
Hello World
Hello World
{[] 0 0}
*/

func Bytes1Buffer2() {
	a := bytes.NewBufferString("Good Night")

	x, err := a.ReadBytes('N') // 读到'N'，包括N
	if err != nil {
		fmt.Println("delim:t err:", err)
	} else {
		fmt.Println(x, string(x))
	}
	fmt.Println(a)

	a.Truncate(0) // 数字0 表示什么都不保留
	fmt.Println(a, a.Len())
	a.WriteString("Good Night")
	fmt.Println(a, a.Len())
	a.Truncate(6) // 数字6 表示保留6个bytes
	fmt.Println(a, a.Len())

	y, err := a.ReadString('d') // 读到'd'，包括d
	if err != nil {
		fmt.Println("delim:N err:", err)
	} else {
		fmt.Println(y)
	}
}

/*
[71 111 111 100 32 78] Good N
ight
 0
Good Night 10
Good N 6
Good
*/
