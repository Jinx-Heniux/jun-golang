package io

import (
	"fmt"
	"strings"
)

func ReadAtExample() {
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 10)
	n, err := reader.ReadAt(p, 4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}
