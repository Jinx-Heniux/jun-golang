package expression

import "fmt"

type Data struct{}

func (Data) TestValue() {
	fmt.Println("In TestValue()")
}

func (*Data) TestPointer() {
	fmt.Println("In TestPointer()")
}

func ExpressionExample4() {
	var p *Data = nil
	p.TestPointer()

	(*Data)(nil).TestPointer() // method value
	(*Data).TestPointer(nil)   // method expression

	// p.TestValue()            // invalid memory address or nil pointer dereference

	// (Data)(nil).TestValue()  // cannot convert nil to type Data
	// Data.TestValue(nil)      // cannot use nil as type Data in function argument
}

/*
In TestPointer()
In TestPointer()
In TestPointer()
*/
