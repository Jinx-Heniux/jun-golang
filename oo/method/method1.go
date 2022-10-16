package method

import (
	"fmt"
)

//结构体
type User struct {
	Name  string
	Email string
}

//方法
func (u User) Notify() {
	fmt.Printf("in Notify: %v, %v, %T, %p \n", u.Name, u.Email, u, &u)
	u.Name = "Jun"
	fmt.Printf("in Notify: %v, %v, %T, %p \n", u.Name, u.Email, u, &u)
}

func MethodExample1() {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	fmt.Printf("u1: %v, %v, %T, %p \n", u1.Name, u1.Email, u1, &u1)
	u1.Notify()
	fmt.Printf("u1: %v, %v, %T, %p \n", u1.Name, u1.Email, u1, &u1)

	// 指针类型调用方法
	fmt.Println()
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	fmt.Printf("u2: %v, %v, %T, %p \n", u2.Name, u2.Email, u2, &u2)
	fmt.Printf("u3: %v, %v, %T, %p,%p\n", u3.Name, u3.Email, u3, u3, &u3)
	u3.Notify()
	fmt.Printf("u2: %v, %v, %T, %p \n", u2.Name, u2.Email, u2, &u2)
	fmt.Printf("u3: %v, %v, %T, %p,%p\n", u3.Name, u3.Email, u3, u3, &u3)
}

/*
u1: golang, golang@golang.com, oo.User, 0xc00002a040
in Notify: golang, golang@golang.com, oo.User, 0xc00002a080
in Notify: Jun, golang@golang.com, oo.User, 0xc00002a080
u1: golang, golang@golang.com, oo.User, 0xc00002a040

u2: go, go@go.com, oo.User, 0xc00002a100
u3: go, go@go.com, *oo.User, 0xc00002a100,0xc00000e030
in Notify: go, go@go.com, oo.User, 0xc00002a140
in Notify: Jun, go@go.com, oo.User, 0xc00002a140
u2: go, go@go.com, oo.User, 0xc00002a100
u3: go, go@go.com, *oo.User, 0xc00002a100,0xc00000e030
*/

// 在Notify方法中，u只是User的副本，即使在方法内被修改了，原本的User不会被更改。
