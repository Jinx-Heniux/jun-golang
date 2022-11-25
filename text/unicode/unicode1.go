package unicode

import (
	"fmt"
	"unicode/utf8"
)

func Unicode1Example1() {
	word := []byte("界")
	fmt.Printf("word: %v, %#v, %p, %p, %d, %d\n", word, word, &word, word, len(word), cap(word))
	for k, v := range word {
		fmt.Printf("k: %d | v: %d, %v, %#v, %T\n", k, v, v, v, v)
	}

	fmt.Println()
	fmt.Println(utf8.Valid(word[:2]))
	fmt.Println(utf8.ValidRune('界'))
	fmt.Println(utf8.ValidString("世界"))

	fmt.Println()
	fmt.Println(utf8.RuneLen('界'))
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('1'))

	fmt.Println()
	fmt.Println(utf8.RuneCount(word))
	fmt.Println(utf8.RuneCountInString("世界"))

	fmt.Println()
	p := make([]byte, 3)
	fmt.Printf("p: %v, %#v, %p, %p, %d, %d\n", p, p, &word, p, len(p), cap(p))
	for k, v := range p {
		fmt.Printf("k: %d | v: %d, %v, %#v, %T\n", k, v, v, v, v)
	}

	fmt.Println()
	utf8.EncodeRune(p, '好')
	fmt.Println(p)

	dr, drs := utf8.DecodeRune(p)
	// fmt.Println(utf8.DecodeRune(p))
	fmt.Printf("drs: %d | dr: %c, %d, %v, %#v, %T\n", drs, dr, dr, dr, dr, dr)

	dris, driss := utf8.DecodeRuneInString("你好马")
	// fmt.Println(utf8.DecodeRuneInString("你好马"))
	fmt.Printf("driss: %d | dris: %c, %d, %v, %#v, %T\n", driss, dris, dris, dris, dris, dris)

	dlr, dlrs := utf8.DecodeLastRune([]byte("你好马"))
	// fmt.Println(utf8.DecodeLastRune([]byte("你好马")))
	fmt.Printf("dlrs: %d | dlr: %c, %d, %v, %#v, %T\n", dlrs, dlr, dlr, dlr, dlr, dlr)

	dlris, dlriss := utf8.DecodeLastRuneInString("你好马")
	// fmt.Println(utf8.DecodeLastRuneInString("你好马"))
	fmt.Printf("dlriss: %d | dlr: %c, %d, %v, %#v, %T\n", dlriss, dlris, dlris, dlris, dlris, dlris)

	fmt.Println()
	fmt.Println(utf8.FullRune(word[:2]))
	fmt.Println(utf8.FullRuneInString("你好"))

	fmt.Println(utf8.RuneStart(word[1]))
	fmt.Println(utf8.RuneStart(word[0]))
}

/*
word: [231 149 140], []byte{0xe7, 0x95, 0x8c}, 0xc0000ac018, 0xc0000ba000, 3, 3
k: 0 | v: 231, 231, 0xe7, uint8
k: 1 | v: 149, 149, 0x95, uint8
k: 2 | v: 140, 140, 0x8c, uint8

false
true
true

3
1
1

1
2

p: [0 0 0], []byte{0x0, 0x0, 0x0}, 0xc0000ac018, 0xc0000ba020, 3, 3
k: 0 | v: 0, 0, 0x0, uint8
k: 1 | v: 0, 0, 0x0, uint8
k: 2 | v: 0, 0, 0x0, uint8

[229 165 189]
drs: 3 | dr: 好, 22909, 22909, 22909, int32
driss: 3 | dris: 你, 20320, 20320, 20320, int32
dlrs: 3 | dlr: 马, 39532, 39532, 39532, int32
dlriss: 3 | dlr: 马, 39532, 39532, 39532, int32

false
true
false
true
*/
