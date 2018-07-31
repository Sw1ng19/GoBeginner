package main
import "os"
import "fmt"

func main() {
	//Go 语言系统编程
	//获取输入内容, 返回值 args 为数组
	args := os.Args[1:]

	fmt.Println("Hello ", args)
}
