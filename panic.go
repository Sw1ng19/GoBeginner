package main
import "os"

func main() {
	panic("慌得一比！")

	_,err := os.Create("/tmp/file")
	if err!=nil {
		panic(err)
	}
}
