package main

import (
	"fmt"
	"bufio"
	"net"
	"log"
	"crypto/tls"
)

func main() {
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		} 

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		msg, err := r.ReadString("\n")
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(msg)

		n, err := conn.Write([]byte("jude\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
