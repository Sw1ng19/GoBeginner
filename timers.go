package main
import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	//在定时器失效以前将一直阻塞
	<-timer1.C
	fmt.Println("timer 1 expired")
}
