package bufio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BufioExample2() {

	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
	fmt.Printf("%v\n", text)
}
