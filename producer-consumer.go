package main

import (
	"fmt"
	"time"
)

func producer(name string, ch chan int) {
	for i:=0; i<4; i++ {
		fmt.Println(name, "product: ", i)

		//写通道
		ch <- i
	}
}

func consumer(name string, ch chan int) {
	for i := range ch {
		fmt.Println(name, "consumer: ", i)
	}
}

func main() {
	ch := make(chan int, 10)
	defer close(ch)

	go producer("生产者A", ch)
	go producer("生产者B", ch)
	go consumer("消费者A", ch)
	go consumer("消费者B", ch)

	time.Sleep(10*time.Second)
}

