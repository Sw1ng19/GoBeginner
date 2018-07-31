package main
import "fmt"

func main() {
	msg := make(chan string, 2)

	//写数据
	go func() {
		msg <- "ping"
		msg <- "buffered"
	}()

	m := <- msg
	fmt.Println(m)
	
	n := <- msg
	fmt.Println(n)
}
