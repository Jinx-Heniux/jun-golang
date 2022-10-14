package fmt

import (
	"fmt"
)

func Placeholder1() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
	/*
		100
		false
		{枯藤}
		struct { name string }{name:"枯藤"}
		struct { name string }
		100%
	*/

	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	/*
		1000001
		A
		65
		101
		41
		41
	*/

	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	/*
		%b	无小数部分、二进制指数的科学计数法，如-123456p-78
		%e	科学计数法，如-1234.456e+78
		%E	科学计数法，如-1234.456E+78
		%f	有小数部分但无指数部分，如123.456
		%F	等价于%f
		%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	*/
	/*
		6946802425218990p-49
		1.234000e+01
		1.234000E+01
		12.340000
		12.34
		12.34
	*/

	s := "枯藤"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
	/*
		枯藤
		"枯藤"
		e69eafe897a4
		E69EAFE897A4
	*/

	a := 18
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
	/*
		0xc0000bc050
		c0000bc050
	*/

	n2 := 88.88
	fmt.Printf("%f\n", n2)
	fmt.Printf("%9f\n", n2)
	fmt.Printf("%.2f\n", n2)
	fmt.Printf("%9.2f\n", n2)
	fmt.Printf("%9.f\n", n2)
	/*
		88.880000
		88.880000
		88.88
		    88.88
		       89
	*/
	/*
		%f	默认宽度，默认精度
		%9f	宽度9，默认精度
		%.2f	默认宽度，精度2
		%9.2f	宽度9，精度2
		%9.f	宽度9，精度0
	*/

	s2 := "枯藤"
	fmt.Printf("%s\n", s2)
	fmt.Printf("%5s\n", s2)
	fmt.Printf("%-5s\n", s2)
	fmt.Printf("%5.7s\n", s2)
	fmt.Printf("%-5.7s\n", s2)
	fmt.Printf("%5.2s\n", s2)
	fmt.Printf("%05s\n", s2)
	/*
		枯藤
		   枯藤
		枯藤
		   枯藤
		枯藤
		   枯藤
		000枯藤
	*/
}
