package oo

import (
	"fmt"
)

//结构体
type User2 struct {
	Name  string
	Email string
}

//方法
func (u *User2) Notify2() {
	fmt.Printf("in Notify: %v, %v, %T, %p, %p \n", u.Name, u.Email, u, &u, u)
	u.Name = "Jun"
	fmt.Printf("in Notify: %v, %v, %T, %p, %p \n", u.Name, u.Email, u, &u, u)
}

func MethodExample2() {
	// 值类型调用方法
	u1 := User2{"golang", "golang@golang.com"}
	fmt.Printf("u1: %v, %v, %T, %p \n", u1.Name, u1.Email, u1, &u1)
	u1.Notify2()
	fmt.Printf("u1: %v, %v, %T, %p \n", u1.Name, u1.Email, u1, &u1)

	// 指针类型调用方法
	fmt.Println()
	u2 := User2{"go", "go@go.com"}
	u3 := &u2
	fmt.Printf("u2: %v, %v, %T, %p \n", u2.Name, u2.Email, u2, &u2)
	fmt.Printf("u3: %v, %v, %T, %p,%p\n", u3.Name, u3.Email, u3, u3, &u3)
	u3.Notify2()
	fmt.Printf("u2: %v, %v, %T, %p \n", u2.Name, u2.Email, u2, &u2)
	fmt.Printf("u3: %v, %v, %T, %p,%p\n", u3.Name, u3.Email, u3, u3, &u3)
}

/*
u1: golang, golang@golang.com, oo.User2, 0xc00002a040
in Notify: golang, golang@golang.com, *oo.User2, 0xc00000e030, 0xc00002a040
in Notify: Jun, golang@golang.com, *oo.User2, 0xc00000e030, 0xc00002a040
u1: Jun, golang@golang.com, oo.User2, 0xc00002a040

u2: go, go@go.com, oo.User2, 0xc00002a0a0
u3: go, go@go.com, *oo.User2, 0xc00002a0a0,0xc00000e038
in Notify: go, go@go.com, *oo.User2, 0xc00000e040, 0xc00002a0a0
in Notify: Jun, go@go.com, *oo.User2, 0xc00000e040, 0xc00002a0a0
u2: Jun, go@go.com, oo.User2, 0xc00002a0a0
u3: Jun, go@go.com, *oo.User2, 0xc00002a0a0,0xc00000e038
*/
