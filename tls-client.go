package main

import (
	"crypto/tls"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	n, err := conn.Write([]byte("hi\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	fmt.Println(string(buf[:n]))
}
