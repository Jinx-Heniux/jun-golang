package expression

import "fmt"

type User1 struct {
	id   int
	name string
}

func (u *User1) Test1() {
	fmt.Printf("%p, %v\n", u, u)
}

func ExpressionExample1() {
	u := User1{1, "Tom"}
	u.Test1()

	mValue := u.Test1
	fmt.Printf("mValue: %p,%T\n", mValue, mValue)
	mValue() // 隐式传递 receiver

	mExpression := (*User1).Test1
	fmt.Printf("mExpression: %p,%T\n", mExpression, mExpression)
	mExpression(&u) // 显式传递 receiver
}

/*
0xc0000ac018, &{1 Tom}
mValue: 0x47f760,func()
0xc0000ac018, &{1 Tom}
mExpression: 0x47f460,func(*main.User)
0xc0000ac018, &{1 Tom}
*/

func (u User1) Test2() {
	fmt.Println(u)
}

func ExpressionExample2() {
	u := User1{1, "Tom"}
	mValue := u.Test2 // 立即复制 receiver，因为不是指针类型，不受后续修改影响。
	fmt.Printf("mValue: %p,%T\n", mValue, mValue)

	u.id, u.name = 2, "Jack"
	u.Test2()

	mValue()
}

/*
mValue: 0x47f920,func()
{2 Jack}
{1 Tom}
*/

func (u *User1) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", u, u)
}

func (u User1) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &u, u)
}

func ExpressionExample3() {
	u := User1{1, "Tom"}
	fmt.Printf("User u: %p, %v\n", &u, u)
	// u2 := User{2, "Tom2"}
	// fmt.Printf("User u2: %p, %v\n", &u2, u2)

	mv := User1.TestValue
	fmt.Printf("mv: %p, %T\n", &mv, mv)
	mv(u)

	mp := (*User1).TestPointer
	fmt.Printf("mp: %p, %T\n", &mp, mp)
	mp(&u)

	mp2 := (*User1).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	fmt.Printf("mp2: %p, %T\n", &mp2, mp2)
	mp2(&u)
}

/*
User u: 0xc00000c030, {1 Tom}
mv: 0xc00000e030, func(main.User)
TestValue: 0xc00000c060, {1 Tom}
mp: 0xc00000e038, func(*main.User)
TestPointer: 0xc00000c030, &{1 Tom}
mp2: 0xc00000e040, func(*main.User)
TestValue: 0xc00000c0a8, {1 Tom}
*/
