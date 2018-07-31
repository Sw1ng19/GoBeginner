package main
import "fmt"
import "time"

func parseStr(from string) {
    for i:=0; i<3; i++ {
        fmt.Println(from,":",i)
    }
}

func main() {
    parseStr("direct")

    //运行协程
    go parseStr("goroutine")
    time.Sleep(time.Second*1)

    fmt.Println("parsing done!")
}
