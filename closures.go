package main
import "fmt"

//函数闭包
func intSeq() func() int {
    i := 0

    return func() int {
        i+=1
        return i
    }
}
