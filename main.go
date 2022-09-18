package main

// learn_function
import (
	"fmt"

	_ "github.com/Jinx-Heniux/jun-golang-hello-world/learn_function"
	_ "github.com/Jinx-Heniux/jun-golang-hello-world/learn_function/lib3"
)

func init() {
	fmt.Println("Init function in package main!")
}

func main() {
	name := "Golang"
	fmt.Printf("Hello, %s!", name)

}

/*
Init function in package learn_function/lib2!
Init function in package learn_function/lib1!
Init function in package learn_function!
Init function in package learn_function/lib3! - Copy
Init function in package learn_function/lib3!
Init function in package learn_function1!
Init function in package main!
Hello, Golang!
*/
