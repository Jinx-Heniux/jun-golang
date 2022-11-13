package flowcontrol

import "fmt"

func Switch1Case1() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 50

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

func Switch1Type1() {
	var x interface{}

	// x = 1
	// x = 1.1
	// x = func(int) float64 { return 1.1 }
	// x = true
	// x = "string"

	//写法一：
	switch i := x.(type) { // 带初始化语句
	case nil:
		fmt.Printf(" x 的类型 :%T\r\n", i)
	case int:
		fmt.Printf("x 是 int 型\n")
	case float64:
		fmt.Printf("x 是 float64 型\n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型\n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型\n")
	default:
		fmt.Printf("未知型\n")
	}

	//写法二
	var j = 0
	switch j {
	case 0:
	case 1:
		fmt.Println("j: 1")
	case 2:
		fmt.Println("j: 2")
	default:
		fmt.Println("j: def")
	}
	//写法三
	var k = 0
	switch k {
	case 0:
		println("k: fallthrough")
		fallthrough
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		fmt.Println("k: 1")
	case 2:
		fmt.Println("k: 2")
	default:
		fmt.Println("k: def")
	}
	//写法三
	var m = 0
	switch m {
	case 0, 1:
		fmt.Println("m: 1")
	case 2:
		fmt.Println("m: 2")
	default:
		fmt.Println("m: def")
	}
	//写法四
	var n = 0
	switch { //省略条件表达式，可当 if...else if...else
	case n > 0 && n < 10:
		fmt.Println("n: i > 0 and i < 10")
	case n > 10 && n < 20:
		fmt.Println("n: i > 10 and i < 20")
	default:
		fmt.Println("n: def")
	}
}
