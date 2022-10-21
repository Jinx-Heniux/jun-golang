package fmt

import (
	"fmt"
	"os"
)

func Fmt1Fprintf1() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./io/fmt/fmtexample1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错, err:", err)
		return
	}
	name := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

func Fmt1Sscan1() {
	var (
		name string
		age  int
	)
	// n, _ := fmt.Sscan("polaris 28", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	n, _ := fmt.Sscan("polaris\n28", &name, &age)
	fmt.Println(n, name, age)
}

func Fmt1Sscanf1() {
	var (
		name string
		age  int
	)
	// n, _ := fmt.Sscanf("polaris 28", "%s%d", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	n, _ := fmt.Sscanf("polaris\n28", "%s%d", &name, &age)
	fmt.Println(n, name, age)
}
