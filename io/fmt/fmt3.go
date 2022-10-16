package fmt

import (
	"bytes"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (person *Person) String() string {
	// 拼接字符串
	buffer := bytes.NewBufferString("This is ")
	buffer.WriteString(person.Name + ", ")
	if person.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}

	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(person.Age))
	buffer.WriteString(" years old.")
	return buffer.String()
}

func (person *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(person.String()))
		f.Write([]byte(" Person has three fields."))
	} else {
		// 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
		f.Write([]byte(fmt.Sprintln(person.String())))
	}
}

func StringerExample1() {
	person1 := &Person{
		"jun",
		18,
		0,
	}
	fmt.Println(person1)
	fmt.Printf("%L\n", person1)
}
