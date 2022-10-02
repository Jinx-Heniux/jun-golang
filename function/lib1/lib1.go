package lib1

import (
	"fmt"

	_ "github.com/Jinx-Heniux/jun-golang/function/lib2"
)

func init() {
	fmt.Println("Init function in package function/lib1!")
}
