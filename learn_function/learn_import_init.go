package learn_function

import (
	"fmt"

	_ "github.com/Jinx-Heniux/jun-golang-hello-world/learn_function/lib1"
	_ "github.com/Jinx-Heniux/jun-golang-hello-world/learn_function/lib2"
)

func init() {
	fmt.Println("Init function in package learn_function!")
}
