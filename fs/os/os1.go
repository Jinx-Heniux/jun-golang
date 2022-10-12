package os

import (
	"fmt"
	"io"
	"os"
)

func OsExample1() {
	var buf [16]byte
	os.Stdin.Read(buf[:])
	os.Stdin.WriteString(string(buf[:]))

	// 输入什么输出什么
}

func OsExample2() {
	// 新建文件
	file, err := os.Create("./fs/os/osexample2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}

func OsExample3() {
	// 打开文件
	file, err := os.Open("./fs/os/osexample2.txt")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

func OsExample4() {
	// 打开源文件
	srcFile, err := os.Open("./fs/os/osexample2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件
	dstFile, err2 := os.Create("./fs/os/osexample4.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 缓冲读取
	buf := make([]byte, 1024)
	for {
		// 从源文件读数据
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//写出去
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()
}
