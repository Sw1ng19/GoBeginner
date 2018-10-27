package main

import (
	"fmt"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc chan string  //读取通道
	wc chan string  //写入通道

	read Reader
	write Writer
}

type ReadFromFile struct {
	path string
}

func (r *ReadFromFile) Read(rc chan string) { //实现 Reader 接口
	line := "message"
	rc <- line
}

func (l *LogProcess) Process() {
	//解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

type WriteToInfluxDB struct {
	influxDBDsn string
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}


func main() {
	//实例化
	r := &ReadFromFile{path: "/tmp/access.log",}
	w := &WriteToInfluxDB{influxDBDsn: "username&password...",}

	lp := &LogProcess{
		rc: make(chan string),
		wc: make(chan string),

		//注入
		read: r,
		write: w,
	}

	//并发执行
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1*time.Second)
}