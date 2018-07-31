package main
import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	//信道干涸，range 等不到信道关闭是不会停止的
	//使用 break 语句主动跳出
	for elem := range queue {
		fmt.Println(elem)
		if len(queue) == 0 {
			break
		}
	}
}
