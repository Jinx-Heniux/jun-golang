package main

/*
import (
	"fmt"

	"github.com/Jinx-Heniux/jun-golang/oo"
)

func main() {
	// var peo oo.People = oo.Student{}
	var peo oo.People = &oo.Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
*/

/*
main.Student
cannot use (Student literal) (value of type Student)
as People value in variable declaration:
Student does not implement People (method Speak has pointer receiver)

修改1
func (stu Student) Speak(think string) (talk string) {

修改2
var peo People = &Student{}

说明 只是*Student类型（不是Student类型）实现了People接口。
*/
