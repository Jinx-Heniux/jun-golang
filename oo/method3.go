package oo

import "fmt"

type Data struct {
	x int
}

func (data Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &data)
}

func (data *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", data)
}

func MethodExample3() {
	d := Data{x: 1}
	p := &d
	fmt.Printf("Data: %v | %p\n", d, p)

	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)

	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
}

/*
Data: {0} | 0xc0000ba000
Value: 0xc0000ba020
Pointer: 0xc0000ba000
Value: 0xc0000ba028
Pointer: 0xc0000ba000
*/
