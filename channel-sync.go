package main
import "fmt"
import "time"

func worker(word chan string) {
	//sleep 时间单位是纳秒
	time.Sleep(5e9)

	word <- "hello" 
}

func main() {
	word := make(chan string, 1)
	go worker(word)

	msg := <- word
	fmt.Println(msg)
}
