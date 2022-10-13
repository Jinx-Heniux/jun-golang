package struct1

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func JsonExample1() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i), // %02d: d表示输出数字，02表示两位，位数不够补零
			Gender: "男",
			ID:     i,
		}
		fmt.Printf("%p\n", c.Students)
		c.Students = append(c.Students, stu) // c.Students的地址不变
		fmt.Printf("%p\n", c.Students)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	fmt.Println()

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

/*
json:{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}

&struct_.Class{Title:"101", Students:[]*struct_.Student{(*struct_.Student)(0xc0000a0690), (*struct_.Student)(0xc0000a06c0), (*struct_.Student)(0xc0000a06f0), (*struct_.Student)(0xc0000a0720), (*struct_.Student)(0xc0000a0780), (*struct_.Student)(0xc0000a07b0), (*struct_.Student)(0xc0000a07e0), (*struct_.Student)(0xc0000a0810), (*struct_.Student)(0xc0000a0840), (*struct_.Student)(0xc0000a0870)}}

*/
