package struct1

import "fmt"

type student3 struct {
	id   int
	name string
	age  int
}

func MapExample1() {
	ce := make(map[int]student3)
	ce[1] = student3{1, "xiaolizi", 22}
	ce[2] = student3{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}

/*
map[1:{1 xiaolizi 22} 2:{2 wang 23}]
map[1:{1 xiaolizi 22}]
*/
