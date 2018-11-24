package main

import (
	"flag"
	"fmt"
)

func main() {
	//.Int() .String() 等返回指针
	/**
	 * compile: go build flags.go
	 * run: ./flags -name=world
	 */
	username := flag.String("name", "", "Please input your username")
	flag.Parse()
	fmt.Println("hello,", *username)
}

