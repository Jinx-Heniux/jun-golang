package io

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func CopyExample() {
	io.Copy(os.Stdout, os.Stdin)
	fmt.Println("Got EOF - bye!")
}

func CopyNExample() {
	io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
}
