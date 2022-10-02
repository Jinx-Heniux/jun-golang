package function

import (
	"fmt"

	_ "github.com/Jinx-Heniux/jun-golang/function/lib1"
	_ "github.com/Jinx-Heniux/jun-golang/function/lib2"
)

func init() {
	fmt.Println("Init function in package function!")
}
