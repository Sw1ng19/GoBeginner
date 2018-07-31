package main
import "fmt"
import "time"

func main() {
	msg := make(chan string)

	go func(word string) {
		msg <- word 
	}("hello")

	time.Sleep(time.Second*1)
	//非阻塞通道选择器
	select {
		case m := <- msg:
			fmt.Println(m)
		default:
			fmt.Println("nothing")
	}
}
