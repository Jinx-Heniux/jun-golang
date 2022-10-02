package io

import (
	"bufio"
	"os"
)

func ReadFromExample() {
	file, err := os.Open("./io/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()

}
