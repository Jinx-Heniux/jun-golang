package io

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func Io2ReadAt1() {
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 10)
	n, err := reader.ReadAt(p, 4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}

func Io2WriteAt1() {
	file, err := os.Create("./io/io/Io2WriteAt1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("Golang中文社区——这里是多余")

	n, err := file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func Io2ReadFrom1() {
	file, err := os.Open("./io/io/Io2WriteAt1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()

}

func Io2WriteTo1() {
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}

func Io2Seeker1() {
	reader := strings.NewReader("Go语言中文网")
	reader.Seek(-6, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}

func Io2ByteRWer1() {
FOREND:
	for {
		fmt.Println("请输入要通过WriteByte写入的一个ASCII字符（b：返回上级；q：退出）：")
		var ch byte
		fmt.Scanf("%c\n", &ch)
		switch ch {
		case 'b':
			fmt.Println("返回上级菜单！")
			break FOREND
		case 'q':
			fmt.Println("程序退出！")
			os.Exit(0)
		default:
			buffer := new(bytes.Buffer)
			err := buffer.WriteByte(ch)
			if err == nil {
				fmt.Println("写入一个字节成功！准备读取该字节……")
				newCh, _ := buffer.ReadByte()
				fmt.Printf("读取的字节：%c\n", newCh)
			} else {
				fmt.Println("写入错误")
			}
		}

	}
}

func Io2LimitedReader1() {
	content := "This Is LimitReader Example"
	reader := strings.NewReader(content)
	limitReader := &io.LimitedReader{R: reader, N: 8}
	for limitReader.N > 0 {
		tmp := make([]byte, 2)
		limitReader.Read(tmp)
		fmt.Printf("%s\n", tmp)
	}

}

func Io2Copy1() {
	io.Copy(os.Stdout, os.Stdin)
	fmt.Println("Got EOF - bye!")
}

func Io2CopyN1() {
	io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
}

func Io2MultiWriter1() {
	file, err := os.Create("./io/io/Io2MultiWriter1.txt")
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

func Io2MultiReader1() {
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
