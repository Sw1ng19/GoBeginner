package main
import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)
	for i:=1; i<5; i++ {
		requests<-i
	}
	//关闭通道后仍可读取数据，否则通道会一直阻塞等待消息
	close(requests)

	limit := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limit
		fmt.Println("request", req, time.Now())
	}
}
