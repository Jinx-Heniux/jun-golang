package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func MultiWriterExample() {
	file, err := os.Create("./io/tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte("Go语言中文网"))
}

func MultiReaderExample() {
	readers := []io.Reader{
		strings.NewReader("from strings reader"),
		bytes.NewBufferString("from bytes buffer"),
	}
	reader := io.MultiReader(readers...) // 可变参数 不定长参数
	data := make([]byte, 0, 128)
	buf := make([]byte, 10)

	for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
		if err != nil {
			panic(err)
		}
		data = append(data, buf[:n]...)
	}
	fmt.Printf("%s\n", data)
}
