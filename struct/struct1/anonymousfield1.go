package struct1

import "fmt"

type User struct {
	id   int
	name string
}

type Manager1 struct {
	User
}

type Manager2 struct {
	User
	title string
}

func (user *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", user, user)
}

func (manager *Manager2) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", manager, manager)
}

func AnonymousFieldExample1() {
	m := Manager1{User{1, "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}

/*
Manager: 0xc00000c030
User: 0xc00000c030, &{1 Tom}
*/

func AnonymousFieldExample2() {
	m := Manager2{User{1, "Tom"}, "Administrator"}

	fmt.Println(m.ToString())

	fmt.Println(m.User.ToString())
}

/*
Manager: 0xc0000a0150, &{{1 Tom} Administrator}
User: 0xc0000a0150, &{1 Tom}
*/
