package main

import "fmt"

func ping(pings chan<-string, msg string) {
	//接收通道
	pings <- msg
}

func pong(pings <-chan string, pongs chan<-string) {
	//pings 发送通道
	msg := <-pings

	//pongs 接收通道
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	
	ping(pings, "hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
