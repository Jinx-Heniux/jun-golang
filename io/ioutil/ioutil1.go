package ioutil

import (
	"fmt"
	"io/ioutil"
)

func wr() {
	err := ioutil.WriteFile("./io/ioutil/ioutilexample1.txt", []byte("www.5lmh.com"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func re() {
	content, err := ioutil.ReadFile("./io/ioutil/ioutilexample1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func Ioutil1Example1() {
	wr()
	re()
}
