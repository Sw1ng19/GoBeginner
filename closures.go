package main
import "fmt"

//函数闭包
func intSeq() func() int {
    fmt.Println("outter func")
    i := 2

    return func() int {
        fmt.Println("inner func")
        i+=1
        return i
    }
}

func main() {
    nextNum := intSeq()
    fmt.Println("result: ", nextNum())
}
