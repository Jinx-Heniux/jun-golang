package io

import (
	"bytes"
	"os"
)

func WriteToExample() {
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}
