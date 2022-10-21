package bufio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Bufio3NewReaderReadSlice() {

	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. a\nIt is the home of gophers"))

	line, _ := reader.ReadSlice('\n')

	fmt.Printf("the line:%s | %p\n", line, line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作

	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s | %p\n", line, line)
	fmt.Printf("the line:%s | %p\n", n, n)
	fmt.Println(string(n))
}

/*
the line:http://studygolang.com.
 | 0xc000108000
the line:It is the home of gophers | 0xc000108000
the line:It is the home of gophers | 0xc000108000
It is the home of gophers
*/

func Bufio3NewReaderSizeReadSlice() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	line, err := reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
	line, err = reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
}

/*
line:http://studygola	error:bufio: buffer full
line:ng.com	error:EOF
*/

func Bufio3NewReaderReadBytes() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))

	line, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作

	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

/*
the line:http://studygolang.com.

the line:http://studygolang.com.

It is the home of gophers

*/

func Bufio3NewReaderSizePeek() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go Peek(reader)
	time.Sleep(time.Second * 1)
	go reader.ReadBytes('\t')
	time.Sleep(5 * time.Second)
}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	time.Sleep(time.Second * 3)
	fmt.Printf("%s\n", line)
}

/*
http://studygo
ng.com.	 It is
*/

func Bufio3NewScannerScan1() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func Bufio3NewScannerSplit() {
	const input = "This is The Golang Standard Library.\nWelcome you!"

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}

/*
8
*/

func Bufio3NewScannerScan2() {
	scanner := bufio.NewScanner(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	if scanner.Scan() {
		scanner.Scan()
		fmt.Printf("%s", scanner.Text())
	}
}

/*
It is the home of gophers
*/

func Bufio3NewScannerScan3() {
	file, err := os.Create("./io/bufio/Bufio3NewScannerScan3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")

	// 将文件 offset 设置到文件开头
	// file.Seek(0, os.SEEK_SET)
	file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
