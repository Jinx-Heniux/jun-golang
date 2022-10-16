package interface1

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func JustifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		// fmt.Printf("x is a string，value is %v\n", v)
		fmt.Printf("x is a %T，value is %v\n", v, v)
	case int:
		// fmt.Printf("x is a int is %v\n", v)
		fmt.Printf("x is a %T，value is %v\n", v, v)
	case bool:
		// fmt.Printf("x is a bool is %v\n", v)
		fmt.Printf("x is a %T，value is %v\n", v, v)
	default:
		fmt.Println("unsupport type！")
	}
}
